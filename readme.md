#  FOOTBALL API

## Overview
ระบบจัดการฟุตบอลพัฒนาโดย GO(Golang) นี้ออกแบบมาเพื่อจัดการและจำลองระบบลีกฟุตบอล ระบบทำงานร่วมกับ API-FOOTBALL (API ภายนอก) เพื่อดึงข้อมูลคะแนน สถิติ ผู้เล่น การแข่งขัน และ ผลการแข่งขัน Feature ได้แก่การตรวจสอบผู้ใช้งาน การสร้างฤดูกาลใหม่ การจัดการผู้เล่น ตารางการแข่งขัน และอัปเดตคะแนนแบบ real-time

## Features
- User Authentication: ลงทะเบียนและเข้าสู่ระบบของผู้ใช้, ใช้ JWT สำหรับการจัดการ sessions และ bcrypt สำหรับการเข้ารหัสรหัสผ่าน เพื่อรักษาความปลอดภัยในการเข้าถึง.
- Season Management: สามารถสร้างและจัดการฤดูกาลใหม่.
- Standings Management: จัดการและอัปเดตตารางคะแนนลีก, ผู้ใช้สามารถเห็นอันดับและสถิติของทีมต่างๆ.
- Player Management: จัดการข้อมูลผู้เล่น, การเพิ่มผู้เล่นใหม่และการอัปเดตสถิติผู้เล่นโดยการดึงข้อมูลจาก API-FOOTBALL (API ภายนอก).
- Match Management: ช่วยให้สามารถสร้างและจัดการการแข่งขัน.
- Live Scores: ฟีเจอร์สำหรับดึงข้อมูลคะแนนสดและอัปเดตคะแนนการแข่งขันแบบ Real-Time จาก API-FOOTBALL (API ภายนอก).
- Match and Player Statistics: แสดงสถิติการแข่งขันและข้อมูลผู้เล่น, เช่น จำนวนประตู, การช่วยทำประตู, การ์ดที่ได้รับ, และเวลาเล่น.

## Technologies Used
- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT, bcrypt
- **External APIs**: API-FOOTBALL - สำหรับดึงข้อมูลการแข่งขันแบบ real-time, ลีค, ทีม, และ ผู้เล่น
- **Containerization**: Docker - สำหรับการรัน PostgreSQL และ pgAdmin4 ใน containers

## Installation
1. Clone the repository
```bash
git clone https://github.com/PlumeSC/API_GO
cd api
```
2. ติดตั้ง Go
3. ติดตั้ง PostgreSQL
4. ตั้งค่า database
5. ตั้งค่า environment variables
    - PORT: พอร์ตที่ API รัน
    - SECRET_KEY: คีย์สำหรับ JWT
    - HOST: โฮสต์ของ database
    - DB_PORT: พอร์ตของ database
    - DB_USER: ชื่อผู้ใช้ database
    - DB_PASSWORD: รหัสผ่าน database
    - DB_NAME: ชื่อ database
    - API_FOOTBALL: API key ของ API-FOOTBALL
    - API_HOST: โฮสต์ของ API-FOOTBALL

## API Endpoints
- **Authentication**
    - 'POST /register' : ลงทะเบียนผู้ใช้
        - Body : '{"firstname":"John", "lastname":"Doe", "email":"JD@mail.com","username":"user", "password":"pass", team:"Brighton"}'
    - 'POST /login' : เข้าสู่ระบบ
        - BODY : '{"username":"user", "password":"pass"}'
- **Season Management**
    - 'POST /createseason' : สร้างฤดูกาลใหม่, ลีค(หากยังไม่มีข้อมูล) และ ตารางคะแนน (ต้องการ authentication และ เป็นAdmin)
        - Body : '{"league":39, "season":2023}'
- **Standings Management**
    - 'GET' /standings : ดึงข้อมูลตารางคะแนน
        - Parameters : 'league=39&season=2023'
- **Player Management**
    - 'POST /createplayer' : อัพเดทและสร้าง Player (ต้องการ authentication และ เป็นAdmin)
        - Body : '{"league":39, "season":2023}'
    - 'GET' /player : ดึงข้อมูลผู้เล่น
        - Parameters : 'playername=Andreas Pereira'
- **Match Management**
    - 'POST /creatematch' : สร้างแมทช์การแข่งขัน (ต้องการ authentication และ เป็นAdmin)
        - Body : '{"league":39, "season":2023}' 
    - 'GET /matches' : ดึงข้อมูลการแข่งขัน หากใส่parameter round จะทำการดึงข้อมูลการแข่งขันนั้นในฤดูกาลนั้นทั้งหมด หากใส่ Parameter team จะทำการดึงการแข่งขันของทีมนั้นในฤดูกาลนั้นทั้งหมด
        - Parameters : 'league=39&season=2023&team=Brighton&round=1'

###  หมายเหตุ
**league**
- Premier League : 39,
- bundesliga : 79,
- Serie A : 135,
- La Liga : 140
                  