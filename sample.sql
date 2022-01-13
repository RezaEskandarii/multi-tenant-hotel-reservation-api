-------------- get rate code prices for reservation
select parent.name as rate_code_name,
       details.rate_code_id,
       details.created_at,
       details.room_id,
       details.id,
       details.date_start,
       details.date_end,
       prices.price,
       prices.guest_count
from rate_code_details details
         join rate_code_detail_prices prices
              on prices.rate_code_detail_id = details.id
         join rate_codes parent on details.rate_code_id = parent.id
where details.room_id = 1
  and prices.guest_count = 1
  and details.min_nights >= 1
  and details.max_nights <= 10
  and details.date_start >= '2019-02-02'
  and details.date_end <= '2025-01-01'

