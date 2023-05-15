
```bash
task build-capsule-http
task build-hello-world
task serve-hello-world
task query-hello-world

task stress-average-hello-world
task stress-high-hello-world

task kill-hello-world
```

```text
Summary:
  Total:        0.1918 secs
  Slowest:      0.1200 secs
  Fastest:      0.0013 secs
  Average:      0.0438 secs
  Requests/sec: 1564.0663
  
  Total data:   15900 bytes
  Size/request: 53 bytes

Response time histogram:
  0.001 [1]     |■
  0.013 [21]    |■■■■■■■■■■■■■
  0.025 [51]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.037 [52]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.049 [65]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.061 [46]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.073 [26]    |■■■■■■■■■■■■■■■■
  0.084 [14]    |■■■■■■■■■
  0.096 [18]    |■■■■■■■■■■■
  0.108 [2]     |■
  0.120 [4]     |■■


Latency distribution:
  10% in 0.0165 secs
  25% in 0.0272 secs
  50% in 0.0393 secs
  75% in 0.0563 secs
  90% in 0.0807 secs
  95% in 0.0891 secs
  99% in 0.1195 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0016 secs, 0.0013 secs, 0.1200 secs
  DNS-lookup:   0.0038 secs, 0.0000 secs, 0.0384 secs
  req write:    0.0017 secs, 0.0000 secs, 0.0157 secs
  resp wait:    0.0351 secs, 0.0012 secs, 0.1200 secs
  resp read:    0.0003 secs, 0.0000 secs, 0.0113 secs

Status code distribution:
  [200] 300 responses
```

