# MonkeyAlerts

## Purpose
The goal of this project is to watch service. If necessary, send mail to watcher.

## Features
- Receive alert mail task
- Receive alert. If times mets the limit,  send mail
- Periodically check services. if false times mets the limit send mail.

## Creat Table Syntax

```sql
  CREATE TABLE `source` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `source` varchar(64) NOT NULL DEFAULT '',
    `target` varchar(64) NOT NULL DEFAULT '',
    `countLimit` char(1) DEFAULT '0',
    `count` char(11) DEFAULT '0',
    `intervalLimit` int(11) DEFAULT '300',
    `lastAlert` int(11) DEFAULT '0',
    PRIMARY KEY (`id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
```