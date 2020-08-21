# gif-progress
[![CircleCI](https://circleci.com/gh/nwtgck/gif-progress.svg?style=shield)](https://circleci.com/gh/nwtgck/gif-progress)

Attach progress bar to animated GIF

![Ray Tracing](doc_assets/ray_tracing.gif)  
The original gif is from <https://github.com/nwtgck/ray-tracing-iow-rust>.

## Installation
Download binaries from [GitHub Release](https://github.com/nwtgck/gif-progress/releases).

## Usage

Generate a gif with bar. 
```bash
cat in.gif | gif-progress > out.gif
```

## Help

```txt
Attach progress bar to animated GIF

Usage:
  gif-progress [flags]

Flags:
      --bar-color string   Bar color (default "#ccc")
      --bar-height int     Bar height (default 5)
      --bar-top            Bar is on top
  -h, --help               help for gif-progress
```

## iOS Version
<img src="https://raw.githubusercontent.com/aheze/ProgressGif/main/Assets/GitHub/Logo/LogoWithShadow.png" width="200">

A similar, [open-source](https://github.com/aheze/ProgressGif) app called ProgressGif is [available on the App Store](https://apps.apple.com/us/app/id1526969349)! Contributions are welcome.
