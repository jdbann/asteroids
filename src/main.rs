use avian2d::prelude::*;
use bevy::prelude::*;
use bevy_rand::prelude::*;
use rand::Rng;

fn main() -> AppExit {
    App::new()
        .add_plugins((
            DefaultPlugins,
            EntropyPlugin::<bevy_prng::WyRand>::with_seed(0_u64.to_ne_bytes()),
            PhysicsPlugins::default(),
            PhysicsDebugPlugin,
        ))
        .insert_resource(Gravity::ZERO)
        .add_systems(Startup, (spawn_camera, spawn_asteroid))
        .run()
}

fn spawn_camera(mut commands: Commands) {
    commands.spawn(Camera2d);
}

#[derive(Component)]
struct Asteroid;

fn spawn_asteroid(
    mut commands: Commands,
    mut rng: Single<&mut bevy_prng::WyRand, With<GlobalRng>>,
) {
    let (vertices, indices) = generate_asteroid_geometry(&mut rng);

    commands.spawn((
        Asteroid,
        Transform::default(),
        RigidBody::Dynamic,
        Collider::convex_decomposition(vertices, indices),
        LinearVelocity(vec2(20.0, 20.0)),
        AngularVelocity(0.5),
    ));
}

fn generate_asteroid_geometry(rng: &mut bevy_prng::WyRand) -> (Vec<Vec2>, Vec<[u32; 2]>) {
    let num_vertices = rng.random_range(8..=16);

    let mut vertices: Vec<Vec2> = Vec::with_capacity(num_vertices);
    let mut indices: Vec<[u32; 2]> = Vec::with_capacity(num_vertices);
    for i in 0..num_vertices {
        vertices.push(
            Vec2::from_angle((360.0 * i as f32 / num_vertices as f32).to_radians()).rotate(Vec2::Y)
                * 50.0
                * rng.random_range(0.75..=1.25),
        );
        indices.push([i as u32, ((i + 1) % num_vertices) as u32]);
    }

    (vertices, indices)
}
