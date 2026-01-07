{
  description = "Development control tool for managing development environments";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    systems.url = "github:nix-systems/default";

    flake-parts = {
      url = "github:hercules-ci/flake-parts";
      inputs.nixpkgs-lib.follows = "nixpkgs";
    };

    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import inputs.systems;
      imports = [
        inputs.treefmt-nix.flakeModule
        inputs.flake-parts.flakeModules.easyOverlay
      ];

      perSystem =
        { inputs', pkgs, ... }:
        let
          devctl = inputs'.gomod2nix.legacyPackages.buildGoApplication {
            pname = "devctl";
            version = "v0.3.0";
            src = ./.;
            modules = ./gomod2nix.toml;

            nativeBuildInputs = with pkgs; [
              git
              ginkgo
            ];

            checkPhase = ''
              ginkgo run --label-filter=!E2E -r .
            '';
          };
        in
        {
          overlayAttrs = {
            inherit devctl;
          };

          packages.devctl = devctl;
          packages.default = devctl;

          devShells.default = pkgs.mkShell {
            buildInputs = with pkgs; [
              direnv
              dprint
              git
              gnumake
              go
              inputs'.gomod2nix.packages.default
              nil
              nixfmt-rfc-style
            ];

            DPRINT = pkgs.dprint + "/bin/dprint";
            GO = pkgs.go + "/bin/go";
            GOMOD2NIX = inputs'.gomod2nix.packages.default + "/bin/gomod2nix";
            NIXFMT = pkgs.nixfmt-rfc-style + "/bin/nixfmt";
          };

          treefmt = {
            programs.nixfmt.enable = true;
            programs.gofmt.enable = true;
          };
        };
    };
}
