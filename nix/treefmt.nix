{
  perSystem =
    { pkgs, ... }:
    {
      treefmt = {
        programs.actionlint.enable = true;
        programs.gofmt.enable = true;
        programs.nixfmt.enable = true;

        programs.dprint = {
          enable = true;
          includes = [
            "*.json"
            "*.md"
            "*.yaml"
            "*.yml"
          ];
          excludes = [
            "flake.lock"
            "gomod2nix.toml"
            "**/*-lock.json"
          ];
          settings = {
            json = { };
            markdown = { };
            yaml = { };
            plugins = pkgs.dprint-plugins.getPluginList (
              p: with p; [
                dprint-plugin-json
                dprint-plugin-markdown
                g-plane-pretty_yaml
              ]
            );
          };
        };
      };
    };
}
