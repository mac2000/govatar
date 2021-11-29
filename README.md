# govatar

**govatar** is a gravatar xmlrpc cli client

Main purpose is to make it easy to update gravatar image with a console tool instead of fighting with bash scripts

In my case it is used in my lol commits like:

```bash
name=$(date +%s)
imagesnap -q -w 2 ~/Documents/lolcommits/$name.jpg
cp ~/Documents/lolcommits/$name.jpg ~/Pictures/avatar.jpg
convert ~/Pictures/avatar.jpg -thumbnail 400x400^ -gravity center -extent 400x400 ~/Pictures/avatar.jpg
# curl -X POST slack.com
govatar --username=marchenko.alexandr@gmail.com --password=123 --set --path=/Users/mac/Pictures.jpg
```

Usage examples are described in [help.txt](help.txt), here are few most wanted:

**list user images**

```bash
govatar --username=user@gmail.com --password=123 --userimages --json
```

```json
[
  {
    "Name": "fcfcc7c079fd5bd411e8c501e94329e5",
    "Rating": 0,
    "URL": "http://en.gravatar.com/userimage/4299573/fcfcc7c079fd5bd411e8c501e94329e5.jpg"
  }
]
```

**save image**

```bash
govatar --username=user@gmail.com --password=123 --save --path=/Users/mac/Pictures/avatar.jpg --json
```

```json
{"id": fcfcc7c079fd5bd411e8c501e94329e5}
```

**use image**

```bash
govatar --username=user@gmail.com --password=123 --use --id=--id=9ad7156f6f7ae6430ecf56d6351a21e5 --email=marchenko.alexandr@gmail.com --json
```

```json
{"success": true}
```

# Binaries

Binaries are available in release section, installation is a simple as downloading and putting them somewhere in a path, e.g.:

```bash
sudo wget https://github.com/mac2000/govatar/releases/download/1.0.0/govatar_darwin_arm64 -O /usr/local/bin/govatar
sudo chmod +x /usr/local/bin/govatar
govatar --username=marchenko.alexandr@gmail.com --password=123 --test
```