# Benchmark

## Scores

```json
{"pass":true,"score":20996,"success":19788,"fail":0,"messages":[]}
isucon@ip-192-168-1-20:~$ /home/isucon/private_isu.git/benchmarker/bin/benchmarker -u /home/isucon/private_isu.git/benchmarker/userdata -t http://57.180.133.103
{"pass":true,"score":22417,"success":21451,"fail":0,"messages":[]}
isucon@ip-192-168-1-20:~$ /home/isucon/private_isu.git/benchmarker/bin/benchmarker -u /home/isucon/private_isu.git/benchmarker/userdata -t http://57.180.133.103
{"pass":true,"score":54116,"success":51743,"fail":0,"messages":[]}
isucon@ip-192-168-1-20:~$ /home/isucon/private_isu.git/benchmarker/bin/benchmarker -u /home/isucon/private_isu.git/benchmarker/userdata -t http://57.180.133.103
{"pass":true,"score":32771,"success":31760,"fail":0,"messages":[]}
// ALTER TABLE posts ADD INDEX created_at_idx (created_at);
isucon@ip-192-168-1-20:~$ /home/isucon/private_isu.git/benchmarker/bin/benchmarker -u /home/isucon/private_isu.git/benchmarker/userdata -t http://57.180.133.103
{"pass":true,"score":76680,"success":73190,"fail":0,"messages":[]}
isucon@ip-192-168-1-20:~$ /home/isucon/private_isu.git/benchmarker/bin/benchmarker -u /home/isucon/private_isu.git/benchmarker/userdata -t http://57.180.133.103
{"pass":true,"score":76728,"success":73290,"fail":0,"messages":[]}
// ALTER TABLE comments ADD INDEX user_id_idx (user_id);
isucon@ip-192-168-1-20:~$ /home/isucon/private_isu.git/benchmarker/bin/benchmarker -u /home/isucon/private_isu.git/benchmarker/userdata -t http://57.180.133.103
{"pass":true,"score":82448,"success":78787,"fail":0,"messages":[]}
// ALTER TABLE users ADD INDEX del_flg_idx (del_flg);
isucon@ip-192-168-1-20:~$ /home/isucon/private_isu.git/benchmarker/bin/benchmarker -u /home/isucon/private_isu.git/benchmarker/userdata -t http://57.180.133.103
{"pass":true,"score":82499,"success":78855,"fail":0,"messages":[]}

```

## Access Log

```
+-------+--------+--------------------+-------+-------+-------+---------+
| COUNT | METHOD |        URI         |  MIN  |  AVG  |  MAX  |   SUM   |
+-------+--------+--------------------+-------+-------+-------+---------+
| 1756  | GET    | /                  | 0.018 | 0.169 | 0.358 | 296.174 |
| 370   | GET    | /posts             | 0.042 | 0.281 | 0.438 | 103.883 |
| 280   | GET    | /@\w+              | 0.033 | 0.254 | 0.392 | 71.132  |
| 1867  | GET    | posts/[0-9]+       | 0.001 | 0.027 | 0.185 | 51.004  |
| 1176  | POST   | /login             | 0.001 | 0.038 | 0.162 | 44.732  |
| 137   | POST   | /register          | 0.012 | 0.051 | 0.117 | 7.035   |
| 2927  | GET    | /favicon.ico       | 0.000 | 0.002 | 0.298 | 6.262   |
| 236   | POST   | /                  | 0.001 | 0.020 | 0.132 | 4.760   |
| 458   | GET    | /login             | 0.001 | 0.004 | 0.131 | 1.910   |
| 116   | POST   | /comment           | 0.004 | 0.016 | 0.151 | 1.904   |
| 229   | GET    | /logout            | 0.001 | 0.003 | 0.086 | 0.784   |
| 137   | GET    | /admin/banned      | 0.001 | 0.005 | 0.015 | 0.620   |
| 33384 | GET    | /image/\d+         | 0.000 | 0.000 | 0.071 | 0.482   |
| 1     | GET    | /initialize        | 0.025 | 0.025 | 0.025 | 0.025   |
| 2927  | GET    | /js/timeago.min.js | 0.000 | 0.000 | 0.000 | 0.000   |
| 2927  | GET    | /js/main.js        | 0.000 | 0.000 | 0.000 | 0.000   |
| 2927  | GET    | /css/style.css     | 0.000 | 0.000 | 0.000 | 0.000   |
+-------+--------+--------------------+-------+-------+-------+---------+

+-------+--------+--------------------+-------+-------+-------+---------+
| COUNT | METHOD |        URI         |  MIN  |  AVG  |  MAX  |   SUM   |
+-------+--------+--------------------+-------+-------+-------+---------+
| 2617  | GET    | /                  | 0.017 | 0.103 | 0.265 | 268.299 |
| 550   | GET    | /posts             | 0.021 | 0.185 | 0.328 | 101.785 |
| 316   | GET    | /@\w+              | 0.056 | 0.260 | 0.462 | 82.310  |
| 1719  | POST   | /login             | 0.001 | 0.032 | 0.175 | 55.744  |
| 2134  | GET    | posts/[0-9]+       | 0.001 | 0.024 | 0.274 | 51.155  |
| 207   | POST   | /register          | 0.012 | 0.046 | 0.149 | 9.485   |
| 3729  | GET    | /favicon.ico       | 0.000 | 0.002 | 0.261 | 6.178   |
| 344   | POST   | /                  | 0.001 | 0.018 | 0.261 | 6.145   |
| 688   | GET    | /login             | 0.001 | 0.004 | 0.167 | 2.430   |
| 136   | POST   | /comment           | 0.003 | 0.015 | 0.261 | 1.978   |
| 344   | GET    | /logout            | 0.001 | 0.003 | 0.021 | 0.955   |
| 207   | GET    | /admin/banned      | 0.001 | 0.004 | 0.022 | 0.882   |
| 49284 | GET    | /image/\d+         | 0.000 | 0.000 | 0.115 | 0.863   |
| 1     | GET    | /initialize        | 0.041 | 0.041 | 0.041 | 0.041   |
| 3729  | GET    | /js/timeago.min.js | 0.000 | 0.000 | 0.000 | 0.000   |
| 3729  | GET    | /js/main.js        | 0.000 | 0.000 | 0.000 | 0.000   |
| 3729  | GET    | /css/style.css     | 0.000 | 0.000 | 0.000 | 0.000   |
+-------+--------+--------------------+-------+-------+-------+---------+



```

## Slow Query
