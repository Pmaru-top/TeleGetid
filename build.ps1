$projectName = "telegetid"
$projectDir = "./main"
$version = "1.0.0"

$targets = @(
    @{os="windows"; arch="386"},
    @{os="windows"; arch="amd64"}
    @{os="windows"; arch="arm"},
    @{os="windows"; arch="arm64"},
    @{os="linux"; arch="386"},
    @{os="linux"; arch="amd64"},
    @{os="linux"; arch="arm"},
    @{os="linux"; arch="arm64"}
)

foreach ($target in $targets) {
    $os = $target.os
    $arch = $target.arch
    $fileName= $projectName + "_" + $os + "_" + $arch

    $env:GOOS = $os
    $env:GOARCH = $arch

    $outputFile = "bin/$fileName"
    if ($os -eq "windows") {
        $outputFile = $outputFile + ".exe"
    }

    $dir = [System.IO.Path]::GetDirectoryName($outputFile)
    if (-not (Test-Path -Path $dir)) {
        New-Item -Path $dir -ItemType Directory
    }

    Write-Host "$os/$arch ..."
    go build -trimpath -ldflags "-s -w -buildid=" -o $outputFile $projectDir

    Write-Host "$os/$arch out: $outputFile"
}
