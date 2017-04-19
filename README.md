# MonkeyAlerts

## Purpose
The goal of this project is to listen and watch service. If necessary, send mail to watcher.

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
    PRIMARY KEY (`id`),
    UNIQUE KEY `sour` (`source`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

  CREATE TABLE `requestTast` (
    `requestTask_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `method` varchar(10) NOT NULL DEFAULT 'GET',
    `URL` varchar(255) NOT NULL DEFAULT '',
    `body` varchar(255) NOT NULL DEFAULT '',
    `headers` varchar(255) NOT NULL DEFAULT '',
    `timeout` int(11) NOT NULL DEFAULT '10',
    `rightStatus` int(11) NOT NULL DEFAULT '200',
    `rightBody` varchar(255) NOT NULL DEFAULT '',
    `source_id` int(11) unsigned NOT NULL,
    PRIMARY KEY (`requestTask_id`),
    KEY `source_id` (`source_id`),
    CONSTRAINT `requesttast_ibfk_1` FOREIGN KEY (`source_id`) REFERENCES `source` (`source_id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
```

## TODO

  - log `alert`
  - log `send mail`
  - load `source`, `requestTask` from Redis