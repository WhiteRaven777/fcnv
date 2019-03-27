# fcnv

fcnv provides Go with fast and easy type conversion.

## Conversion table
<table>
<tr><th>func</th><th>from</th><th>to</th></tr>
<tr><td>Atoi</td><td rowspan="13">string</td><td>int</td></tr>
<tr><td>Atoi8</td><td>int8</td></tr>
<tr><td>Atoi16</td><td>int16</td></tr>
<tr><td>Atoi32</td><td>int32</td></tr>
<tr><td>Atoi64</td><td>int64</td></tr>
<tr><td>Atoui</td><td>uint</td></tr>
<tr><td>Atoui8</td><td>uint8</td></tr>
<tr><td>Atoui16</td><td>uint16</td></tr>
<tr><td>Atoui32</td><td>uint32</td></tr>
<tr><td>Atoui64</td><td>uint64</td></tr>
<tr><td>Atof32</td><td>float32</td></tr>
<tr><td>Atof64</td><td>float64</td></tr>
<tr><td>Atob</td><td>[]byte</td></tr>
<tr><td>Itoa</td><td>int</td><td rowspan="11">string</td></tr>
<tr><td>Itoa8</td><td>int8</td></tr>
<tr><td>Itoa16</td><td>int16</td></tr>
<tr><td>Itoa32</td><td>int32</td></tr>
<tr><td>Itoa64</td><td>int64</td></tr>
<tr><td>Uitoa</td><td>uint</td></tr>
<tr><td>Uitoa8</td><td>uint8</td></tr>
<tr><td>Uitoa16</td><td>uint16</td></tr>
<tr><td>Uitoa32</td><td>uint32</td></tr>
<tr><td>Uitoa64</td><td>uint64</td></tr>
<!--<tr><td>Atof32</td><td>float32</td></tr>-->
<!--<tr><td>Atof64</td><td>float64</td></tr>-->
<tr><td>Btoa</td><td>[]byte</td></tr>
<tr><td>Byte2Int</td><td>[]byte</td><td>int</td></tr>
<tr><td>Int2Byte</td><td>int</td><td>[]byte</td></tr>
<tr><td>Byte2Bool</td><td>[]byte</td><td>bool</td></tr>
<tr><td>Bool2Byte</td><td>bool</td><td>[]byte</td></tr>
<tr><td>Bool2Int</td><td>bool</td><td>int</td></tr>
<tr><td>Int2Bool</td><td>int</td><td>bool</td></tr>
<tr><td>Bool2Uint</td><td>bool</td><td>uint</td></tr>
<tr><td>Uint2Bool</td><td>uint</td><td>bool</td></tr>
<tr><td>Bool2Str</td><td>bool</td><td>string</td></tr>
<tr><td>Str2Bool</td><td>string</td><td>bool</td></tr>
<tr><td>Struct2Json</td><td>interface{}</td><td>string</td></tr>
<tr><td>Datetime2Date</td><td>time.Time</td><td>time.Time</td></tr>
<tr><td>Hankaku2Zenkaku</td><td>string</td><td>string</td></tr>
<tr><td>Zenkaku2Hankaku</td><td>string</td><td>string</td></tr>
</table>

## Benchmark
### Atoi* / Atoui*
* old: BenchmarkAto(|u)i\d?ByParse(I|Ui)nt
* new: BenchmarkNewAto(|u)i\d?

```
name          old time/op     new time/op    delta
Atoi-12       29.5ns ± 4%    22.9ns ± 4%  -22.23%  (p=0.000 n=72+76)
Atoi8-12      26.0ns ± 6%    20.5ns ±10%  -21.22%  (p=0.000 n=78+77)
Atoi16-12     25.3ns ± 5%    20.9ns ± 8%  -17.37%  (p=0.000 n=78+73)
Atoi32-12     25.7ns ± 9%    21.3ns ± 9%  -17.30%  (p=0.000 n=71+68)
Atoi64-12     28.9ns ± 6%    25.0ns ± 9%  -13.30%  (p=0.000 n=75+77)
Atoui-12      37.4ns ± 3%    33.2ns ± 5%  -11.23%  (p=0.000 n=67+70)
Atoui8-12     24.3ns ±12%    23.1ns ±13%   -4.92%  (p=0.000 n=72+76)
Atoui16-12    24.3ns ± 7%    22.8ns ±10%   -6.00%  (p=0.000 n=72+70)
Atoui32-12    23.5ns ± 3%    17.6ns ± 6%  -25.09%  (p=0.000 n=72+73)
Atoui64-12    25.9ns ± 7%    22.4ns ± 7%  -13.55%  (p=0.000 n=72+73)

name          old alloc/op    new alloc/op   delta
Atoi-12        2.00B ± 0%     2.00B ± 0%     ~     (all equal)
Atoi8-12       9.00B ± 0%     9.00B ± 0%     ~     (all equal)
Atoi16-12      6.00B ± 0%     6.00B ± 0%     ~     (all equal)
Atoi32-12      4.00B ± 0%     3.61B ±17%   -9.69%  (p=0.000 n=80+80)
Atoi64-12      2.00B ± 0%     2.00B ± 0%     ~     (all equal)
Atoui-12       8.00B ± 0%     8.00B ± 0%     ~     (all equal)
Atoui8-12      16.0B ± 0%     16.0B ± 0%     ~     (all equal)
Atoui16-12     12.0B ± 0%     12.0B ± 0%     ~     (all equal)
Atoui32-12     7.00B ± 0%     3.00B ± 0%  -57.14%  (p=0.000 n=80+80)
Atoui64-12     4.00B ± 0%     4.00B ± 0%     ~     (all equal)

name          old allocs/op   new allocs/op  delta
Atoi-12          0.00            0.00          ~     (all equal)
Atoi8-12         0.00            0.00          ~     (all equal)
Atoi16-12        0.00            0.00          ~     (all equal)
Atoi32-12        0.00            0.00          ~     (all equal)
Atoi64-12        0.00            0.00          ~     (all equal)
Atoui-12         0.00            0.00          ~     (all equal)
Atoui8-12        0.00            0.00          ~     (all equal)
Atoui16-12       0.00            0.00          ~     (all equal)
Atoui32-12       0.00            0.00          ~     (all equal)
Atoui64-12       0.00            0.00          ~     (all equal)
```

### test
```
$ vgo test -bench Ato -o test.bin -cpuprofile=cpu.prof -benchmem -count 100 -short -timeout 60m | tee run.log
$ benchstat run.log
$ pprof -http=":8888" test.bin cpu.prof
```
