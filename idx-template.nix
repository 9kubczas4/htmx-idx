{ pkgs, backend ? "go", ... }: {
  packages = [];
  bootstrap = ''
    cp -rf ${backend} "$WS_NAME"
    chmod -R +w "$WS_NAME"
    mv "$WS_NAME" "$out"
  '';
}
