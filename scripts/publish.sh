#!/bin/bash
# publish.sh builds and publishes a release.
set -o nounset -o errexit -o pipefail

ROOT=$(dirname $0)/..
PUBLISH=$GOPATH/src/github.com/pulumi/home/scripts/publish.sh
PUBLISH_GOOS=("linux" "windows" "darwin")
PUBLISH_GOARCH=("amd64")
PUBLISH_PROJECT="pulumi-github"

if [ ! -f $PUBLISH ]; then
    >&2 echo "error: Missing publish script at $PUBLISH"
    exit 1
fi

echo "Publishing SDK build to s3://eng.pulumi.com/:"
for OS in "${PUBLISH_GOOS[@]}"
do
    for ARCH in "${PUBLISH_GOARCH[@]}"
    do
        export GOOS=${OS}
        export GOARCH=${ARCH}

        RELEASE_INFO=($($(dirname $0)/make_release.sh))
        ${PUBLISH} ${RELEASE_INFO[0]} "${PUBLISH_PROJECT}/${OS}/${ARCH}" ${RELEASE_INFO[@]:1}
    done
done

echo "Publishing NPM package to NPMjs.com:"
pushd ${ROOT}/pack/bin && \
    npm publish && \
    npm info 2>/dev/null || true && \
    popd
