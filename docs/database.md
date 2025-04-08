# database design

## models

### users

- id: int, primary key, auto increment
- username: string, unique, not null
- password: string, not null
- email: string, unique, not null
- phone: string, unique
- avatar: string, url
- credit_score: int, default 100
- status: uint, default 0 (0: normal, 1: locked, 2: deleted) not null
- last_login_at: datetime, default current timestamp
- created_at: datetime, default current timestamp
- updated_at: datetime, default current timestamp on update current timestamp

### user_addresses

- id: int, primary key, auto increment
- user_id: int, foreign key to users(id), not null
- receiver_name: string, not null
- receiver_phone: string, not null
- province: string, not null
- city: string, not null
- district: string, not null
- detail_address: string, not null
- is_default: tinyint, default 0 (0: no, 1: yes), not null
- created_at: datetime, default current timestamp
- updated_at: datetime, default current timestamp on update current timestamp
- Foreign key (user_id) references users(id) on delete cascade

### categories

- id: int, primary key, auto increment
- parent_id: int, foreign key to categories(id), default 0 (0: root category)
- name: string, not null
- icon: string, not null
- sort_order: int, default 0 comment 'smaller number means higher priority'
- status: uint, default 0 (0: normal, 1: hidden) not null
- created_at: datetime, default current timestamp
- updated_at: datetime, default current timestamp on update current timestamp

### items

- id: int, primary key, auto increment
- seller_id: int, foreign key to users(id), not null
- category_id: int, foreign key to categories(id), not null
- title: string, not null
- description: text, not null
- price: decimal(10, 2), not null
- orginal_price: decimal(10, 2), default 0.00
- condition: uint, default 0 (0: new, 1: almost new, 2: slightly used, 3: used, 4: damaged) not null
- location: string, not null
- status: uint, default 0 (0: normal, 1: sold, 2: down, 3: checking, 4: invalid) not null
- view_count: int, default 0
- favorite_count: int, default 0
- is_free_shipping: tinyint, default 0 (0: no, 1: yes)
- shipping_fee: decimal(10, 2), default 0.00
- created_at: datetime, default current timestamp
- updated_at: datetime, default current timestamp on update current timestamp

### items_images

- id: int, primary key, auto increment
- item_id: int, foreign key to items(id), not null
- image_url: string, not null
- sort_order: int, default 0 comment 'smaller number means higher priority'
- is_cover: tinyint, default 0 (0: no, 1: yes), not null
- created_at: datetime, default current timestamp
- updated_at: datetime, default current timestamp on update current timestamp

### tags

- id: int, primary key, auto increment
- name: string, not null
- created_at: datetime, default current timestamp
- updated_at: datetime, default current timestamp on update current timestamp

### item_tag_relations

- id: int, primary key, auto increment
- item_id: int, foreign key to items(id), not null
- tag_id: int, foreign key to tags(id), not null
- created_at: datetime, default current timestamp
- updated_at: datetime, default current timestamp on update current timestamp
- Unique key (item_id, tag_id) 

### orders

- id: int, primary key, auto increment
- order_no: string, unique, not null
- buyer_id: int, foreign key to users(id), not null
- seller_id: int, foreign key to users(id), not null
- item_id: int, foreign key to items(id), not null
- item_snapshot: text, not null
- price: decimal(10, 2), not null
- shipping_fee: decimal(10, 2), default 0.00
- total_price: decimal(10, 2), not null
- address_id: int, foreign key to user_addresses(id), not null
- address_snapshot: text, not null
- status: uint, default 0 (0: unpaid, 1: paid, 2: shipped, 3: completed, 4: canceled, 5: refunded) not null
- payment_method: uint, default 0 (0: alipay, 1: wechat, 2: bank transfer) not null
- payment_time: datetime, default null
- shipping_time: datetime, default null
- delivery_company: string, default null
- tracking_number: string, default null
- completed_time: datetime, default null
- cancel_reason: string, default null
- created_at: datetime, default current timestamp
- updated_at: datetime, default current timestamp on update current timestamp

### reviews

- id: int, primary key, auto increment
- order_id: int, foreign key to orders(id), not null
- reviewer_id: int, foreign key to users(id), not null
- reviewee_id: int, foreign key to users(id), not null
- item_id: int, foreign key to items(id), not null, 
- rating: int, not null, comment '1-5 stars'
- content: text, not null
- anonymous: tinyint, default 0 (0: no, 1: yes), not null
- created_at: datetime, default current timestamp
- updated_at: datetime, default current timestamp on update current timestamp

### review_replies

- id: int, primary key, auto increment
- review_id: int, foreign key to reviews(id), not null
- user_id: int, foreign key to users(id), not null
- content: text, not null
- created_at: datetime, default current timestamp
- updated_at: datetime, default current timestamp on update current timestamp

### favorites

- id: int, primary key, auto increment
- user_id: int, foreign key to users(id), not null
- item_id: int, foreign key to items(id), not null
- created_at: datetime, default current timestamp
- updated_at: datetime, default current timestamp on update current timestamp
- Unique key (user_id, item_id)

### browsing_histories

### messages

### system_notifications

### wallet_transactions

### reports

### search_histories