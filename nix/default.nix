{ inputs, ... }:
{
  systems = import inputs.systems;

  imports = [
    inputs.treefmt-nix.flakeModule
    inputs.flake-parts.flakeModules.easyOverlay

    ./apps.nix
    ./checks.nix
    ./devshell.nix
    ./packages.nix
    ./treefmt.nix
  ];
}
