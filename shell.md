# leetcode 学习总结

## shell

- 1.Read from the file file.txt and output the tenth line to stdout.  -- https://leetcode.com/problems/tenth-line/description/
```
sed -n '10p' file.txt
```

- 2.Read from the file file.txt and output all valid phone numbers to stdout. -- https://leetcode.com/problems/valid-phone-numbers/description/
```
grep -E '^(\([0-9]{3}\) ){1}[0-9]{3}-[0-9]{4}$|^([0-9]{3}-){2}[0-9]{4}$'  file.txt
```
注意： grep -E 引入或逻辑。

- 3. # Read from the file words.txt and output the word frequency list to stdout.  -- https://leetcode.com/problems/word-frequency/description/
```
sed 's/\\n/ /g' words.txt | awk '{ for(i=1;i<=NF;i++) print $i}' | sort | uniq -c | awk '{print $2,$1}' | sort -nr -k2
```
误区：1. 需要将 ‘\n’转义为空格； 2. sort -n 按照数字排序
注： sort 用法：
```
文本排序：sort
      -n：数值排序
      -r: 降序
      -t: 字段分隔符
      -k: 以哪个字段为关键字进行排序
      -u: 排序后相同的行只显示一次
      -f: 排序时忽略字符大小写
```

http://wiki.ubuntu.org.cn/Shell%E7%BC%96%E7%A8%8B%E5%9F%BA%E7%A1%80#.E4.BB.8E.E7.AC.AC.E4.B8.80.E8.A1.8C.E5.BC.80.E5.A7.8B
