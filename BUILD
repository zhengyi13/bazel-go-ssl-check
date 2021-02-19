load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/zhengyi13/bazel-go-ssl-check",
)

go_library(
    name = "go_default_library",
    srcs = ["ssl_expiry.go"],
    importpath = "github.com/example/project",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_zhengyi13_prober//:go_default_library",
        "@in_gopkg_yaml_v2//:go_default_library",
    ],
)

go_binary(
    name = "ssl_check",
    embed = [":go_default_library"],
    importpath = "github.com/zhengyi13/bazel-go-ssl-check",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["ssl_expiry_test.go"],
    embed = [":go_default_library"],
)
