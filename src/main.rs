use avian2d::prelude::*;
use bevy::prelude::*;
use bevy_rand::prelude::*;
use rand::Rng;

const GRAPPLING_HOOK_RANGE: f32 = 300.0;

fn main() -> AppExit {
    App::new()
        .add_plugins((
            DefaultPlugins,
            EntropyPlugin::<bevy_prng::WyRand>::new(),
            PhysicsPlugins::default(),
            PhysicsDebugPlugin,
        ))
        .insert_gizmo_config(
            PhysicsGizmos {
                shapecast_color: Some(Color::srgba(1.0, 0.0, 0.0, 0.2)),
                ..default()
            },
            GizmoConfig::default(),
        )
        .init_resource::<WorldBounds>()
        .insert_resource(Gravity::ZERO)
        .add_systems(Startup, (spawn_camera, spawn_asteroid, spawn_player))
        .add_systems(
            FixedUpdate,
            (
                move_player,
                fire_cannon,
                fire_grappling_hook,
                apply_wrapping,
                despawn_beyond_world_bounds,
            ),
        )
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
        ShapeCaster::new(Collider::circle(5.0), vec2(0.0, 20.0), 0.0, Dir2::NORTH)
            .with_max_hits(1)
            .with_max_distance(GRAPPLING_HOOK_RANGE),
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

#[derive(Component)]
struct DespawnBeyondWorldBounds;

#[derive(Resource)]
struct WorldBounds {
    half_size: Vec2,
}

impl Default for WorldBounds {
    fn default() -> Self {
        Self {
            half_size: vec2(400.0, 300.0) * 3.0,
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

fn despawn_beyond_world_bounds(
    mut commands: Commands,
    bounds: Res<WorldBounds>,
    query: Query<(Entity, &Transform), With<DespawnBeyondWorldBounds>>,
) {
    let world_rect = Rect::from_center_half_size(Vec2::ZERO, bounds.half_size);
    for (entity, transform) in query {
        if !world_rect.contains(transform.translation.xy()) {
            commands.entity(entity).despawn();
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

#[derive(Component)]
struct Bullet;

fn fire_cannon(
    input: Res<ButtonInput<KeyCode>>,
    mut commands: Commands,
    query: Query<(&Transform, &LinearVelocity), With<Player>>,
) {
    if input.just_pressed(KeyCode::Space) {
        let (transform, linear_velocity) = query.single().unwrap();
        commands.spawn((
            Bullet,
            transform.with_translation(transform.translation + transform.local_y() * 16.0),
            RigidBody::Dynamic,
            Collider::segment(vec2(0.0, 0.0), vec2(0.0, 5.0)),
            MassProperties2d::new(0.1, 0.1, vec2(0.0, 2.5)).to_bundle(),
            LinearVelocity(
                transform.local_y().xy()
                    * (200.0 + linear_velocity.0.dot(transform.local_y().xy())),
            ),
            DespawnBeyondWorldBounds,
        ));
    }
}

fn fire_grappling_hook(
    input: Res<ButtonInput<KeyCode>>,
    mut commands: Commands,
    query: Query<(Entity, &ShapeHits)>,
    transform_query: Query<&GlobalTransform>,
) {
    if input.just_pressed(KeyCode::KeyG) {
        let (entity, shape_hits) = query.single().unwrap();

        let Some(hit) = shape_hits.first() else {
            return;
        };

        let hit_transform = transform_query.get(hit.entity).unwrap();

        commands.spawn(
            DistanceJoint::new(entity, hit.entity)
                .with_local_anchor1(vec2(0.0, -5.0))
                .with_local_anchor2(
                    hit_transform
                        .affine()
                        .inverse()
                        .transform_point3(hit.point1.extend(0.0))
                        .xy(),
                )
                .with_limits(5.0, GRAPPLING_HOOK_RANGE + 30.0),
        );
    }
}
