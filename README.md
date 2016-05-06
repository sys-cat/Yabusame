# Notifity to Something Devices System is "Yabusame"

Android, iOS, Web にデバイスに関わらず同じ内容をPushするシステムです。現在モック作成中。

## 作成予定の機能

* Push内容の作成
* Pushの送信
* ユーザ毎のPush先の変更
	* デバイストークンの管理
	* 2垢対応
* カレンダー機能
* カテゴリ機能
* セグメンテーション機能
* 既読管理
* APNs, GCMの管理

## モック版で作成する機能

* Push内容の作成
* ユーザ毎のPush先の変更
	* デバイストークンの管理
	* 2垢対応
* カレンダー機能
* カテゴリ機能

## DB構成

*Index: Push*

Type : Item
| field | mapping | default term |
| Title | string | "no title" |
| Body | string | "no body" |
| Category | integer | "" |
| CreatedAt | Date | Now |
| UpdatedAt | Date | Now |

Type : User
| field | mapping | default term |
| SystemId | string |  |
| Device | string | "None" |
| Sub | string | "" |
| CreatedAt | Date | Now |
| UpdatedAt | Date | Now |

Type : Calendar
| field | mapping | default term |
| Date | Date |  |
| ItemId | string |  |
| Users | JSON | "[]" |
| Device | string | "None" |
| CreatedAt | Date | Now |
| UpdatedAt | Date | Now |

Type : Category
| field | mapping | default term |
| Name | string | "Dummy Category" |
| CreatedAt | Date | Now |
| UpdatedAt | Date | Now |

GCM, APNsを使える様にするにはAPIKeyが必要なのでモック版では行わないでおく
