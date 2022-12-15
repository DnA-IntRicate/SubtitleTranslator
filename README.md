# SubtitleTranslator

SubtitleTranslator is a barebones command-line subtitle translating application for SRT files written in [Go 1.18](https://go.dev) that utilises [go-googletrans](https://github.com/Conight/go-googletrans) for translations. In-future this application may be extended to be able to arbitrarily translate textfiles of any type.

## Clone from Github

```batch
git clone --recursive https://github.com/DnA-IntRicate/SubtitleTranslator.git
```

## Examples

### Any language to English

```batch
SubtitleTranslator -i InputFile.srt -o OutputFile.srt
```

### Turkish to English

```batch
SubtitleTranslator -i InputFile.srt -o OutputFile.srt -s tr -d en
```

See [Language Codes](https://developers.google.com/admin-sdk/directory/v1/languages) for specifying translation languages.

## Usage

```txt
SubtitleTranslator v1.0.

Valid switches:
-i, --in, --input               Specify the input file path.
-o, --out, --output             Specify the file path to ouput translated file.
-s, --src, --source             Specify the source file's language. (Set to 'auto' by default).
-d, --dst, --destination        Specify the language to translate to. (Set to 'English (en)' by default).
-q, --quiet                     Don't output translation results in terminal.

Valid usages:
Convert from any language implicitly to English: 'SubtitleTranslator -i InputFile.srt -o OutputFile.srt'
Convert explicitly from Turkish implicitly to English: 'SubtitleTranslator -i InputFile.srt -o OutputFile.srt -s tr
Convert explicitly from Turkish implicitly to English: 'SubtitleTranslator -i InputFile.srt -o OutputFile.srt -s tr
Convert explicitly from English explicitly to Urdu: 'SubtitleTranslator -i InputFile.srt -o OutputFile.srt -s en -d ur
```

## Building

### Build

```batch
go build .
OR
go build -ldflags "-s -w"
```

This will output the executable to the project's root.

### Install

```batch
go install .
OR
go install -ldflags "-s -w"
```

This will build the executable to GOPATH.

### Install Go

[Download Go](https://go.dev/dl/)

Debian:
```shell
sudo apt-get update
sudo apt-get install golang-go
```

## License

This application is distributed under the [Apache License Version 2.0](https://www.apache.org/licenses/LICENSE-2.0).

```txt
Copyright 2022 Adam Foflonker

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
