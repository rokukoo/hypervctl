# hyperctl

A simple sdk for Microsoft Hyper-V virtual machine management which is implemented in Go and using wmi.

## Installation

```bash
go get github.com/rokukoo/hyperctl
```

## Features

### NetworkAdapter

| 名称         | 介绍          | 进度 |
|------------|-------------|----|
| List       | 列表查询网络适配器   | ✅  |
| Enable     | 启用网络适配器     | ❌  |
| Disable    | 禁用网络适配器     | ❌  |
| Configure  | 配置网络适配器     | ❌  |
| FindByName | 根据名称查询网络适配器 | ❌  |

### VirtualMachine

- [x] Create
- [x] Destroy
- [x] Delete
- [x] Start
- [x] Stop
- [x] ForceStop
- [x] Shutdown
- [x] Reboot
- [x] ForceReboot
- [x] Pause
- [x] Restore
- [x] Save
- [x] Resume
- [ ] Snapshot
- [ ] ModifyName
- [x] ModifySpecByName
- [ ] ModifyInternalIpv4Address
- [x] GetById
- [x] List
- [x] FindByName
  - [x] FirstByName

### VirtualHardDisk

- [x] Create
- [x] DeleteByPath
- [x] Attach
- [x] Detach
- [x] Resize
- [x] GetByPath

### VirtualSwitch

- [x] Create
    - [x] External
    - [x] Bridge
    - [x] Internal
    - [x] Private
- [x] Delete
- [x] ChangeType
- [ ] EnableVLan/DisableVLan
- [ ] SetVlanId
- [ ] GetById
- [x] FindByName
- [x] List

### VirtualNetworkAdapter

[//]: # (- [x] Create)

[//]: # (- [x] Delete)
- [x] Attach
- [x] Detach
- [x] Connect
- [x] Disconnect
- [ ] SetMacAddress
- [x] SetBandwidth
- [ ] EnableVLan/DisableVLan
    - [ ] SetVLanId
- [x] List
- [x] FindByName

### Cluster

