workspace(
    name = "profefe",
    managed_directories = {"@npm": ["node_modules"]},
)

## General rules
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
#load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

######################
# GOLANG SUPPORT
######################

## rules_go
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "d6b2513456fe2229811da7eb67a444be7785f5323c6708b38d851d2b51e54d83",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.30.0/rules_go-v0.30.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.30.0/rules_go-v0.30.0.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.17.6")

## Gazelle - used for Golang external dependencies
http_archive(
    name = "bazel_gazelle",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//thirdparty/go:repositories.bzl", "go_repositories")

# gazelle:repository_macro thirdparty/go/repositories.bzl%go_repositories
go_repositories()

gazelle_dependencies()
######################
# DOCKER SUPPORT
######################

## rules_docker
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "59536e6ae64359b716ba9c46c39183403b01eabfbd57578e84398b4829ca499a",
    strip_prefix = "rules_docker-0.22.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.22.0/rules_docker-v0.22.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

######################
# PROTOBUF
######################

http_archive(
    name = "com_google_protobuf",
    sha256 = "9111bf0b542b631165fadbd80aa60e7fb25b25311c532139ed2089d76ddf6dd7",
    strip_prefix = "protobuf-3.18.1",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.18.1.tar.gz"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

######################
# SKYLIB
######################
http_archive(
    name = "bazel_skylib",
    sha256 = "c6966ec828da198c5d9adbaa94c05e3a1c7f21bd012a0b29ba8ddbccb2c93b0d",
    urls = [
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.1.1/bazel-skylib-1.1.1.tar.gz",
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.1.1/bazel-skylib-1.1.1.tar.gz",
    ],
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

######################
# DYNAMIC STUFF
######################
load("@io_bazel_rules_docker//container:pull.bzl", "container_pull")

container_pull(
    name = "golang_distroless_arm64",
    digest = "sha256:1a6d571a1247b4ef4dceda8a3919c872eb3f5801cf42858b882ee1a53c8fc6ff",
    registry = "gcr.io",
    repository = "distroless/static",
    tag = "latest",
)

container_pull(
    name = "golang_distroless_amd64",
    digest = "sha256:49f33fac9328ac595cb74bd02e6a186414191c969de0d8be34e6307c185acb8e",
    registry = "gcr.io",
    repository = "distroless/static",
    tag = "latest",
)