-- Clean up existing data to avoid conflicts
TRUNCATE TABLE price_history, subscriptions, payment_methods, categories, users CASCADE;

-- Procedural generation of data
DO $$
DECLARE
    u_id UUID;
    cat_id UUID;
    pm_id UUID;
    sub_id UUID;
    i INT;
    j INT;
    cat_names TEXT[] := ARRAY['Streaming', 'Software', 'Fitness', 'Gaming', 'Education', 'Utilities', 'News', 'Music'];
    cat_name TEXT;
    pm_types TEXT[] := ARRAY['credit_card', 'digital_wallet', 'debit_card', 'bank_transfer'];
    service_names TEXT[] := ARRAY['Netflix', 'Spotify', 'AWS', 'Gym', 'Xbox', 'Adobe', 'YouTube', 'Kindle', 'Apple One', 'Disney+', 'Hulu', 'HBO Max'];
BEGIN
    -- Loop to create 100 Users
    FOR i IN 1..100 LOOP
        INSERT INTO users (id, name, email, password_hash)
        VALUES (
            gen_random_uuid(), 
            'User ' || i, 
            'user' || i || '@example.com', 
            '$2a$10$SomeHashedPasswordForTest' || i
        )
        RETURNING id INTO u_id;

        -- Create a few Categories for this user (plus some random ones)
        FOREACH cat_name IN ARRAY cat_names LOOP
             INSERT INTO categories (id, user_id, name, icon, color)
             VALUES (gen_random_uuid(), u_id, cat_name, 'star', '#000000')
             RETURNING id INTO cat_id;
        END LOOP;

        -- Create a few Payment Methods for this user
        FOR j IN 1..3 LOOP
            INSERT INTO payment_methods (id, user_id, name, type)
            VALUES (
                gen_random_uuid(), 
                u_id, 
                'Method ' || j, 
                pm_types[1 + (j % array_length(pm_types, 1))]
            )
            RETURNING id INTO pm_id;
        END LOOP;

        -- Create 100 Subscriptions for this user
        FOR j IN 1..100 LOOP
            -- Pick a random category (just picking the last one created for simplicity or a random one from DB would be better but expensive inside loop. 
            -- Let's just creat a fresh category or re-use one. For performance, let's pick one we just made.
            -- To make it realistic, let's select one random category ID for this user.
            SELECT id INTO cat_id FROM categories WHERE user_id = u_id ORDER BY random() LIMIT 1;
            
            -- Select random payment method
            SELECT id INTO pm_id FROM payment_methods WHERE user_id = u_id ORDER BY random() LIMIT 1;

            INSERT INTO subscriptions (
                id, user_id, category_id, payment_method_id, service_name, price, currency, cycle, status, next_billing_date
            )
            VALUES (
                gen_random_uuid(),
                u_id,
                cat_id,
                pm_id,
                service_names[1 + (floor(random() * array_length(service_names, 1)))::int],
                (random() * 100)::numeric(10, 2),
                'USD',
                'monthly',
                'active',
                CURRENT_DATE + (floor(random() * 30) || ' days')::interval
            )
            RETURNING id INTO sub_id;

            -- Add some price history
            IF (random() > 0.8) THEN
                INSERT INTO price_history (id, subscription_id, old_price, new_price, reason)
                VALUES (
                    gen_random_uuid(),
                    sub_id,
                    (random() * 50)::numeric(10, 2),
                    (random() * 100)::numeric(10, 2),
                    'Random price adjustment'
                );
            END IF;
        END LOOP;
    END LOOP;
END $$;
