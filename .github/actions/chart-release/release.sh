#!/bin/bash

# Adapted  from https://github.com/helm/chart-releaser-action/blob/master/cr.sh

main() {
    install_chart_releaser

    crversion=v1.1.1
    owner=$(cut -d '/' -f 1 <<< "$GITHUB_REPOSITORY")
    repo=$(cut -d '/' -f 2 <<< "$GITHUB_REPOSITORY")

    sed -i 's/0.0.0/$VERSION/g' charts/tic-tac-toe/Chart.yaml

    cr package charts/tic-tac-toe
    cr upload --owner $owner --git-repo $repo --token $GH_TOKEN
    cr index -i index.yaml --owner $owner --git-repo $repo --charts-repo https://$owner.github.io/$repo

    cat index.yaml
}

install_chart_releaser() {
    if [[ ! -d "$RUNNER_TOOL_CACHE" ]]; then
        echo "Cache directory '$RUNNER_TOOL_CACHE' does not exist" >&2
        exit 1
    fi

    local arch
    arch=$(uname -m)

    local cache_dir="$RUNNER_TOOL_CACHE/ct/$crversion/$arch"
    if [[ ! -d "$cache_dir" ]]; then
        mkdir -p "$cache_dir"

        echo "Installing chart-releaser..."
        curl -sSLo cr.tar.gz "https://github.com/helm/chart-releaser/releases/download/$crversion/chart-releaser_${crversion#v}_linux_amd64.tar.gz"
        tar -xzf cr.tar.gz -C "$cache_dir"
        rm -f cr.tar.gz

        echo 'Adding cr directory to PATH...'
        export PATH="$cache_dir:$PATH"
    fi
}

main