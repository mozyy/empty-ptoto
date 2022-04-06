# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.166.1/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.18, 1.17
ARG VARIANT="1"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Option] Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next line to use go get to install anything else you need
# RUN go get -x <your-dependency-or-tool>

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

RUN apt-get update

# ------------ docker -----------
# https://github.com/microsoft/vscode-dev-containers/blob/main/containers/docker-from-docker/.devcontainer/Dockerfile

# [Option] Install zsh
ARG INSTALL_ZSH="true"
# [Option] Upgrade OS packages to their latest versions
ARG UPGRADE_PACKAGES="false"
# [Option] Enable non-root Docker access in container
ARG ENABLE_NONROOT_DOCKER="true"
# [Option] Use the OSS Moby CLI instead of the licensed Docker CLI
ARG USE_MOBY="true"
# [Option] Select CLI version
ARG CLI_VERSION="latest"

# Enable new "BUILDKIT" mode for Docker CLI
ENV DOCKER_BUILDKIT=1

# Install needed packages and setup non-root user. Use a separate RUN statement to add your
# own dependencies. A user of "automatic" attempts to reuse an user ID if one already exists.
ARG USERNAME=automatic
ARG USER_UID=1000
ARG USER_GID=$USER_UID
COPY library-scripts/*.sh /tmp/library-scripts/
RUN apt-get update \
    && /bin/bash /tmp/library-scripts/common-debian.sh "${INSTALL_ZSH}" "${USERNAME}" "${USER_UID}" "${USER_GID}" "${UPGRADE_PACKAGES}" "true" "true" \
    # Use Docker script from script library to set things up
    && /bin/bash /tmp/library-scripts/docker-debian.sh "${ENABLE_NONROOT_DOCKER}" "/var/run/docker-host.sock" "/var/run/docker.sock" "${USERNAME}" "${USE_MOBY}" "${CLI_VERSION}" \
    # Clean up
    && apt-get autoremove -y && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts/

# Setting the ENTRYPOINT to docker-init.sh will configure non-root access to 
# the Docker socket if "overrideCommand": false is set in devcontainer.json. 
# The script will also execute CMD if you need to alter startup behaviors.
ENTRYPOINT [ "/usr/local/share/docker-init.sh" ]
CMD [ "sleep", "infinity" ]

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# ------------ docker -----------

# --------------- custom ------

ARG PROTOC_VERSION=3.20.0
ARG PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP} && \
    unzip $PROTOC_ZIP -d ~/.local && \
    rm -f $PROTOC_ZIP

ARG PROTOC_WEB_VERSION=1.3.1
ARG PROTOC_WEB_ZIP=protoc-gen-grpc-web-${PROTOC_WEB_VERSION}-linux-x86_64

RUN curl -L -o /usr/local/bin/protoc-gen-grpc-web https://github.com/grpc/grpc-web/releases/download/${PROTOC_WEB_VERSION}/${PROTOC_WEB_ZIP} && \
    chmod +x /usr/local/bin/protoc-gen-grpc-web

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

ARG ANGULAR_CLI_VERSION=latest
# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g @angular/cli@${ANGULAR_CLI_VERSION}" 2>&1

ARG PROTOC_GEN_GO_VERSION=latest
ARG PROTOC_GEN_GO_GRPC_VERSION=latest
ARG PROTOC_GEN_GORM_VERSION=latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@$PROTOC_GEN_GO_VERSION &&\
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$PROTOC_GEN_GO_GRPC_VERSION &&\
    go install github.com/mozyy/protoc-gen-gorm@$PROTOC_GEN_GORM_VERSION
# go install github.com/micro/micro/v3/cmd/protoc-gen-micro && \
# go install github.com/micro/micro/v3/cmd/protoc-gen-openapi

ENV GOCACHE=/workspaces/cache/go_cache
ENV GOMODCACHE=/workspaces/cache/go_mod_cache
ENV YARN_CACHE_FOLDER=/workspaces/cache/yarn_cache