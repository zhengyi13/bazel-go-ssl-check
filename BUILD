load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/example/project",
)

go_library(
    name = "go_default_library",
    srcs = ["ssl_expiry.go"],
    importpath = "github.com/zhengyi13/bazel-go-ssl-check",
    visibility = ["//visibility:private"],
    deps = ["@in_gopkg_yaml_v2//:go_default_library"],
)

go_binary(
    name = "ssl_check",
    embed = [":go_default_library"],
    importpath = "github.com/zhengyi13/bazel-go-ssl-check",
    visibility = ["//visibility:public"],
)
