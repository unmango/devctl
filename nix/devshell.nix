{
  perSystem =
    {
      config,
      inputs',
      pkgs,
      ...
    }:
    {
      devShells.default = pkgs.mkShell {
        packages =
          (with pkgs; [
            git
            gnumake
            go
            goreleaser
            gopls
            nil
            nixfmt
          ])
          ++ [
            inputs'.gomod2nix.legacyPackages.gomod2nix
            config.treefmt.build.wrapper
          ];
      };
    };
}
