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
        .init_resource::<WorldBounds>()
        .insert_resource(Gravity::ZERO)
        .add_systems(Startup, (spawn_camera, spawn_asteroid, spawn_player))
        .add_systems(FixedUpdate, (move_player, apply_wrapping))
        .add_systems(Update, draw_world_bounds)
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
        Transform::from_translation(vec3(
            rng.random_range(-400.0..=400.0),
            rng.random_range(-300.0..=300.0),
            0.0,
        )),
        RigidBody::Dynamic,
        Collider::convex_decomposition(vertices, indices),
        LinearVelocity(rng.random::<Vec2>() * rng.random_range(-20.0..=20.0)),
        AngularVelocity(rng.random::<f32>() * rng.random_range(-1.0..=1.0)),
        Mass(2.0),
        ApplyWrapping,
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

#[derive(Component)]
struct Player;

#[derive(Component)]
struct ShipImpulses {
    thrust: f32,
    turn: f32,
}

fn spawn_player(mut commands: Commands) {
    let (vertices, indices) = generate_player_geometry();

    commands.spawn((
        Player,
        Transform::default(),
        RigidBody::Dynamic,
        Collider::convex_decomposition(vertices, indices),
        Mass(1.0),
        AngularDamping(5.0),
        ShipImpulses {
            thrust: 2.0,
            turn: 20.0,
        },
        ApplyWrapping,
    ));
}

fn generate_player_geometry() -> (Vec<Vec2>, Vec<[u32; 2]>) {
    (
        vec![
            vec2(0.0, 15.0),
            vec2(-10.0, -15.0),
            vec2(0.0, -10.0),
            vec2(10.0, -15.0),
        ],
        vec![[0, 1], [1, 2], [2, 3], [3, 0]],
    )
}

fn move_player(
    input: Res<ButtonInput<KeyCode>>,
    mut query: Query<(&ShipImpulses, Forces), With<Player>>,
) {
    for (ship_impulses, mut forces) in &mut query {
        if input.pressed(KeyCode::ArrowUp) {
            forces.apply_local_linear_impulse(vec2(0.0, ship_impulses.thrust));
        }

        if input.pressed(KeyCode::ArrowLeft) {
            forces.apply_angular_impulse(ship_impulses.turn);
        }

        if input.pressed(KeyCode::ArrowRight) {
            forces.apply_angular_impulse(-1.0 * ship_impulses.turn);
        }
    }
}

#[derive(Component)]
struct ApplyWrapping;

#[derive(Resource)]
struct WorldBounds {
    half_size: Vec2,
}

impl Default for WorldBounds {
    fn default() -> Self {
        Self {
            half_size: vec2(400.0, 300.0),
        }
    }
}

fn apply_wrapping(bounds: Res<WorldBounds>, mut query: Query<&mut Transform, With<ApplyWrapping>>) {
    for mut transform in &mut query {
        if transform.translation.x <= -bounds.half_size.x {
            transform.translation.x += bounds.half_size.x * 2.0;
        } else if transform.translation.x >= bounds.half_size.x {
            transform.translation -= bounds.half_size.x * 2.0;
        }

        if transform.translation.y <= -bounds.half_size.y {
            transform.translation.y += bounds.half_size.y * 2.0;
        } else if transform.translation.y >= bounds.half_size.y {
            transform.translation.y -= bounds.half_size.y * 2.0;
        }
    }
}

fn draw_world_bounds(mut gizmos: Gizmos, bounds: Res<WorldBounds>) {
    gizmos.rect_2d(
        Isometry2d::IDENTITY,
        bounds.half_size * 2.0,
        Color::srgb(1.0, 0.0, 0.0),
    );
}
