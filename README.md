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

Generate a gif with the interactive cli:
- Download the bash file [gif-progress.sh](https://github.com/nwtgck/gif-progress/#)
```bash
./gif-progress.sh

# OR

sh /path/to/gif-progress.sh
```

If you have problems executing the file make sure it's executable using the following command:
```bash
sudo chmod a+x /path/to/gif-progress.sh
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
  
Examples: 
  Change height:
    cat in.gif | gif-progress --bar-height=15 > out.gif
  Change color:
    cat in.gif | gif-progress --bar-color=#07b7ae > out.gif
  Set bar on top:
    cat in.gif | gif-progress --bar-top > out.gif
  Change height and color:
    cat in.gif | gif-progress --bar-color=#07b7ae --bar-height=15 > out.gif
```
