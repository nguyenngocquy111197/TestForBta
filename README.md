I: Hướng dẫn chạy
    Port:
        8080: BookingService
        8081: SendService
        8082: PricingService

    C1: 
    step 1: Bấm vào Makefile
    Step 2: Để chạy được 3 service thay đổi các key-value sau
        MAIN_PATH := ./exec/pricing or ./exec/send  or ./exec/booking 
        DEFAULT_EXEC := ./bin/pricing.exec or ./bin/pricing.send or ./bin/pricing.booking
    Step 3: 
        Thay đổi biến môi trường hệ điều hành:
            GOOS  # linux or windows or darwin 
    Step 4:
        chạy lệnh make start ở terminal 
    

    C2: 
    step 1: Bấm vào Makefile
    Step 2: Để chạy được 3 service thay đổi các key-value sau
        MAIN_PATH := ./exec/pricing or ./exec/send  or ./exec/booking 
        DEFAULT_EXEC := ./bin/pricing.exec or ./bin/pricing.send or ./bin/pricing.booking
    Step 3: 
        Thay đổi biến môi trường hệ điều hành:
            GOOS  # linux or windows or darwin 
    Step 4:
        Chạy lệnh make builder tại terminal 
            - Nó sẽ build thành các file .exec
    Step 5:
        Chạy file .exec
    
    Lưu ý: set up file config tương ứng với mỗi service (tham khảo trong folder config)

I: Mô tả ngắn gọn API
    - Flow BookingService :
        - /account/listRole : danh sách role người dùng có thể tạo 
        - /account/create : tạo tài khoản hàng và người cung cấp dịch vụ:

        - /familyService/insert : tạo dịch vụ
        - /familyService/get : danh sách dịch vụ đã tạo 

        - /booking : tạo đơn 
        - /booking/check/status : kiểm tra trạng thái đơn hàng
        - /booking/update/status : cập nhật trạng thái đơn hàng

    - Flow PrincingService :
        - /calculate: tính toán số tiền dựa trên dịch vụ và ngày

    - Flow SendService :
        - /send/info/service: gửi thông tin đến người cung cấp dịch vụ

Để xem chi tiết hơn về các API , tham khảo ở 2 file Postman và Word đính kèm trong email 