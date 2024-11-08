---
prev:
  text: 'What is kyanos'
  link: './what-is-kyanos'
next: false
---

# Quick Start

## Installation Requirements

- Kernel version 3.10 or higher

## Installation and Running {#prerequire}

You can download a statically linked binary compatible with x86_64 and arm64 architectures from the [release page](https://github.com/hengyoush/kyanos/releases):

```bash
tar xvf kyanos_vx.x.x_linux_x86.tar.gz
```

Then, run:
```bash
kyanos watch 
```

If the following table appears:
![kyanos quick start success](/quickstart-success.png)
🎉 Congratulations! Kyanos has started successfully.

> [!TIP]
> Did the command above fail? No worries—check the FAQ below to see if your situation is covered. If not, feel free to open a [GitHub issue](https://github.com/hengyoush/kyanos/issues)!

## FAQ {#faq}
### No Output After Running kyanos watch

Check if you see a log like this when starting Kyanos:

```js
find btf file exactly failed, try to find a lower version btf file...
```
If so, this might be due to a missing BTF file on your system. You can manually download the BTF file matching your kernel version from [this link](https://mirrors.openanolis.cn/coolbpf/btf/) or [this link](https://github.com/aquasecurity/btfhub-archive/), and then specify the downloaded BTF file using the `--btf` option when starting Kyanos.

## Next Steps
- For a quick guide on using Kyanos, see: [Learn Kyanos in 5 Minutes](./how-to)