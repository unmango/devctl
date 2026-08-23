{
  perSystem =
    { config, ... }:
    {
      checks.unit = config.packages.devctl.overrideAttrs (_: {
        pname = "devctl-unit-tests";

        doCheck = true;
        # `go tool ginkgo` builds the version pinned in go.mod, avoiding the
        # version skew warning that pkgs.ginkgo produces.
        checkPhase = ''
          runHook preCheck
          go tool ginkgo run --race --trace --label-filter='!E2E' -r .
          runHook postCheck
        '';

        installPhase = ''
          runHook preInstall
          touch $out
          runHook postInstall
        '';
      });
    };
}
