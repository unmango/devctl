{
  perSystem =
    { inputs', pkgs, ... }:
    let
      inherit (inputs'.gomod2nix.legacyPackages) gomod2nix;

      # These mutate the working tree and/or need network access, so they are
      # apps rather than derivations. Sandboxed equivalents live in checks.
      goTools = with pkgs; [
        git
        go
      ];
    in
    {
      apps.tidy = {
        meta.description = "Sync go.sum and gomod2nix.toml with go.mod";
        program = pkgs.writeShellApplication {
          name = "tidy";
          runtimeInputs = goTools ++ [ gomod2nix ];
          text = ''
            go mod tidy
            gomod2nix
          '';
        };
      };

      apps.test = {
        meta.description = "Run the unit test suites";
        program = pkgs.writeShellApplication {
          name = "test";
          runtimeInputs = goTools;
          text = ''
            go tool ginkgo run --label-filter='!E2E' -r "''${@:-.}"
          '';
        };
      };

      apps.e2e = {
        meta.description = "Run the end-to-end test suite (needs network)";
        program = pkgs.writeShellApplication {
          name = "e2e";
          runtimeInputs = goTools;
          text = ''
            go tool ginkgo run --label-filter=E2E -r "''${@:-.}"
          '';
        };
      };
    };
}
