据[奶粉佬](https://nf.asoul-rec.com)所说：~~鼠历~~05:00(CST)分割天压力较小

### Global
- 二进制编码为小端
- 文字编码原样复制

#### Danmaku 弹幕流
- 单文件 最大4G 不切割
- 当弹幕到来时:
- 空闲时每`time.Hour*2`关闭一次
- 直播时每`time.Hour/4`关闭一次
- 关闭时向TagList写入检查点

#### CheckPoint 检查点
- 单文件 最大4G 不切割 超出覆写最旧数据
- 在弹幕流中关闭时写入
- 格式: Tag Length Value
- 获取直播间信息，与前一时刻比较，增量写入
- 获取主播信息，同上
- 生成时间、弹幕流偏移（压缩后）
- 写入时间、直播间信息、主播信息
- `Ctrl+C`处理：不获取信息，仅生成时间戳并保存结束弹幕数据流

#### CPList 标记的索引 每月压缩
- 单文件 不切割 不限大小 与CheckPoint大小有关
- Encoding: binary.LittleEndian
- UnixMilli uint64 开始时间
- TagList   uint32 偏移    
- Extra     uint32 额外数据  


对应关系： 
    一个用户对应一个直播间对应一个TagList对应一个弹幕流文件
    每个弹幕流文件最大4G-1，超过覆盖最早的部分
