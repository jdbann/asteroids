use avian2d::prelude::*;
use bevy::prelude::*;

fn main() -> AppExit {
    App::new()
        .add_plugins((
            DefaultPlugins,
            PhysicsPlugins::default(),
            PhysicsDebugPlugin,
        ))
        .add_systems(Startup, (spawn_camera, spawn_asteroid))
        .run()
}

fn spawn_camera(mut commands: Commands) {
    commands.spawn(Camera2d);
}

#[derive(Component)]
struct Asteroid;

fn spawn_asteroid(mut commands: Commands) {
    commands.spawn((Asteroid, Collider::circle(50.0)));
}
