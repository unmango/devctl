{
  perSystem =
    {
      inputs',
      lib,
      pkgs,
      ...
    }:
    let
      inherit (inputs'.gomod2nix.legacyPackages) buildGoApplication;

      devctl = buildGoApplication {
        pname = "devctl";
        # release-please owns this file; it is the single source of truth for
        # the released version.
        version = (lib.importJSON ../.release-please-manifest.json).".";
        src = ../.;
        modules = ../gomod2nix.toml;

        nativeBuildInputs = with pkgs; [ git ];

        # Tests live in checks.unit so `nix build` stays fast and a failing
        # test is attributable to the check rather than the package.
        doCheck = false;
      };
    in
    {
      overlayAttrs = {
        inherit devctl;
      };

      packages.devctl = devctl;
      packages.default = devctl;
    };
}
