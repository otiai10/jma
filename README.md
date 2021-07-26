# Yet another weather forecast API client in Japan

[![Go](https://github.com/otiai10/jma/actions/workflows/go.yaml/badge.svg)](https://github.com/otiai10/jma/actions/workflows/go.yaml)
[![codecov](https://codecov.io/gh/otiai10/jma/branch/main/graph/badge.svg?token=yeK0wKMzm7)](https://codecov.io/gh/otiai10/jma)

# Client SDK

```go
import "github.com/otiai10/jma/api"

client := &api.Client{}
res, err := client.Forecast(jma.Fukuoka)
```

# CLI

```zsh
# Install
% go install github.com/otiai10/jma/cli/forecast@latest

# Usage
% forecast fukuoka
		07/27	07/28	07/29	07/30	07/31	08/01	08/02
【福岡県】	☀️  	🌤  20	🌥  30	🌤  30	🌤  20	🌤  20	🌤  20
```

# Refs

* https://qiita.com/e_toyoda/items/7a293313a725c4d306c0
* https://qiita.com/michan06/items/48503631dd30275288f7
* https://www.jma.go.jp/jma/kishou/know/amedas/ame_master.pdf
* https://forest.watch.impress.co.jp/docs/serial/yajiuma/1309318.html
* https://www.jma.go.jp/bosai/forecast/data/forecast/130000.json
* https://www.jma.go.jp/bosai/forecast/data/overview_forecast/130000.json?__time__=202102240506
* `Forecast.Const.TELOPS` https://www.jma.go.jp/bosai/map.html
* https://www.jma.go.jp/bosai/forecast/img/100.svg
