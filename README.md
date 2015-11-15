# plugin-cf-manifest-generator

Create a manifest file from 0.

Currently Cloud Foundry Cli can only generate a manifest file from an existing app. This plugin help you to create one without using an existing app.

Manifest-generator will ask you some questions to generate manifest. It will also find data from your Cloud Foundry instance to generate it.

You can see it in action here:
![manifest-generator](/img/manifest-generator.gif "Demo manifest-generator")

## Installation from cf cli (prefered)

```bash
$ cf add-plugin-repo CF-Community http://plugins.cloudfoundry.org/
$ cf install-plugin manifest-generator -r CF-Community
```

## Installation from release binaries

#### On OSX using release binaries

```bash
$ wget -O $GOPATH/bin/manifest-generator https://github.com/ArthurHlt/plugin-cf-manifest-generator/releases/download/v1.0.0/manifest-generator_darwin_amd64
$ cf install-plugin $GOPATH/bin/manifest-generator
```

#### On Linux using release binaries

64bit:

```bash
$ wget -O $GOPATH/bin/manifest-generator https://github.com/ArthurHlt/plugin-cf-manifest-generator/releases/download/v1.0.0/manifest-generator_linux_amd64
$ cf install-plugin $GOPATH/bin/manifest-generator
```

32bit:

```bash
$ wget -O $GOPATH/bin/manifest-generator https://github.com/ArthurHlt/plugin-cf-manifest-generator/releases/download/v1.0.0/manifest-generator_linux_386
$ cf install-plugin $GOPATH/bin/manifest-generator
```

#### On Windows using release binaries

64bit:

```bash
#in your browser download https://github.com/ArthurHlt/plugin-cf-manifest-generator/releases/download/v1.0.0/manifest-generator_windows_amd64.exe and place it in $GOPATH/bin/manifest-generator.exe

$ cf install-plugin $env:GOPATH/bin/manifest-generator.exe
```

32bit:

```bash
#in your browser download https://github.com/ArthurHlt/plugin-cf-manifest-generator/releases/download/v1.0.0/manifest-generator_windows_386.exe and place it in $GOPATH/bin/manifest-generator.exe

$ cf install-plugin $env:GOPATH/bin/manifest-generator.exe
```

## installation using go get

#### On *nix using go get

```bash
$ go get github.com/ArthurHlt/plugin-cf-manifest-generator
$ cf install-plugin $GOPATH/bin/plugin-cf-manifest-generator
```

#### On Windows using go get

```bash
$ go get github.com/ArthurHlt/plugin-cf-manifest-generator
$ cf install-plugin $GOPATH/bin/plugin-cf-manifest-generator.exe
```

## Usage

```bash
cf manifest-generator [-n path/or/file/name/of/your/future/manifest.yml]
```
