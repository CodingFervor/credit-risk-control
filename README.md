# Financial Credit Risk Control System | 金融级信贷风控系统

[English](#english) | [中文](#中文)

## English
Financial-grade credit risk control platform with loan application, multi-dimensional rule engine (Drools-style), credit scoring, repayment management, overdue collection, blacklist, and real-time risk dashboard.

### Features
- Loan application & approval workflow
- Configurable risk rule engine with decision tree scoring
- Credit limit management & dynamic adjustment
- Repayment scheduling & collection automation
- Blacklist & fraud detection
- Real-time risk monitoring dashboard (Flink streaming)

### Tech Stack: Go 1.22 + Gin + PostgreSQL + Redis + Kafka + Drools (via gRPC) + Flink

## 中文
金融级信贷风控决策平台，支持授信申请、多维度风控规则引擎、评分卡、额度管理、还款催收、黑名单管理和实时风控大屏。

### 快速开始
```bash
git clone https://github.com/CodingFervor/credit-risk-control.git
cd credit-risk-control && docker-compose up -d && go run cmd/api/main.go
```
## License: MIT
