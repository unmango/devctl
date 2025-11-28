{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    gomod2nix.url = "github:nix-community/gomod2nix";
    gomod2nix.inputs.nixpkgs.follows = "nixpkgs";
    treefmt-nix.url = "github:numtide/treefmt-nix";
    treefmt-nix.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs =
    inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [ inputs.treefmt-nix.flakeModule ];

      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
      perSystem =
        { inputs', pkgs, ... }:
        {
          devShells.default = pkgs.mkShell {
            buildInputs = with pkgs; [
              dprint
              git
              gnumake
              go
              # inputs'.gomod2nix.packages.gomod2nix
              nixfmt-rfc-style
            ];
          };

          treefmt = {
            programs.nixfmt.enable = true;
            programs.gofmt.enable = true;
          };
        };
    };
}
