-- 1. Limpa dados existentes para evitar conflitos de chave
TRUNCATE TABLE price_history, subscriptions, payment_methods, categories, users CASCADE;

-- 2. Procedural generation of data
DO $$
DECLARE
    u_id UUID;
    cat_id UUID;
    pm_id UUID;
    sub_id UUID;
    i INT;
    j INT;
    
    -- 🎯 CORREÇÃO: Variável TEXT para iterar sobre nomes de categorias
    current_cat_name TEXT; 
    
    -- Dados Globais Fixo
    global_cat_names TEXT[] := ARRAY['Streaming', 'Software', 'Fitness', 'Gaming', 'Education', 'Utilities', 'News', 'Music'];
    
    -- Arrays para IDs globais
    category_ids UUID[] := ARRAY[]::UUID[];
    payment_method_ids UUID[] := ARRAY[]::UUID[];

    -- Dados Variáveis para Assinaturas
    service_names TEXT[] := ARRAY['Netflix', 'Spotify', 'AWS', 'Gym', 'Xbox', 'Adobe', 'YouTube', 'Kindle', 'Apple One', 'Disney+', 'Hulu', 'HBO Max'];
    pm_global_names TEXT[] := ARRAY['Cartão de Crédito', 'PIX', 'Boleto', 'Débito Automático'];
    pm_global_types TEXT[] := ARRAY['CREDIT_CARD', 'PIX', 'BANK_SLIP', 'DEBIT_CARD'];

BEGIN
    -- 3. Inserir CATEGORIAS GLOBAIS (executado apenas uma vez)
    -- ✅ Usando a variável TEXT (current_cat_name) para iterar sobre o array TEXT[]
    FOREACH current_cat_name IN ARRAY global_cat_names LOOP
         INSERT INTO categories (id, name, icon, color)
         VALUES (
            gen_random_uuid(), 
            current_cat_name, -- Uso da variável TEXT
            'icon_' || lower(current_cat_name), 
            '#' || substring(md5(random()::text) FROM 1 FOR 6) 
         )
         RETURNING id INTO cat_id;
         category_ids := array_append(category_ids, cat_id);
    END LOOP;

    -- 4. Inserir MÉTODOS DE PAGAMENTO GLOBAIS (executado apenas uma vez)
    FOR i IN 1..array_length(pm_global_names, 1) LOOP
        INSERT INTO payment_methods (id, name, type)
        VALUES (
            gen_random_uuid(), 
            pm_global_names[i], 
            pm_global_types[i]
        )
        RETURNING id INTO pm_id;
        payment_method_ids := array_append(payment_method_ids, pm_id);
    END LOOP;

    -- 5. Loop principal para criar 100 Usuários e suas Assinaturas
    FOR i IN 1..100 LOOP
        INSERT INTO users (id, name, email, password_hash)
        VALUES (
            gen_random_uuid(), 
            'User ' || i, 
            'user' || i || '@example.com', 
            '$2a$10$SomeHashedPasswordForTest' || i
        )
        RETURNING id INTO u_id;

        -- 6. Cria 5 a 15 Assinaturas para cada usuário
        FOR j IN 1..(5 + floor(random() * 10)::int) LOOP
            
            -- 🎯 Seleciona um ID de Categoria Global aleatório
            cat_id := category_ids[1 + (floor(random() * array_length(category_ids, 1)))::int];
            
            -- 🎯 Seleciona um ID de Método de Pagamento Global aleatório
            pm_id := payment_method_ids[1 + (floor(random() * array_length(payment_method_ids, 1)))::int];

            INSERT INTO subscriptions (
                id, user_id, category_id, payment_method_id, service_name, price, currency, cycle, status, next_billing_date
            )
            VALUES (
                gen_random_uuid(),
                u_id,
                cat_id, -- FK Global
                pm_id,  -- FK Global
                service_names[1 + (floor(random() * array_length(service_names, 1)))::int],
                (random() * 100)::numeric(10, 2),
                'BRL', 
                CASE floor(random() * 3)
                    WHEN 0 THEN 'monthly'
                    WHEN 1 THEN 'quarterly'
                    ELSE 'annually'
                END,
                'active',
                CURRENT_DATE + (floor(random() * 30) || ' days')::interval
            )
            RETURNING id INTO sub_id;

            -- 7. Adiciona histórico de preços (com menor probabilidade)
            IF (random() > 0.9) THEN
                INSERT INTO price_history (id, subscription_id, old_price, new_price, reason)
                VALUES (
                    gen_random_uuid(),
                    sub_id,
                    (random() * 50)::numeric(10, 2),
                    (random() * 100)::numeric(10, 2),
                    'Ajuste aleatório de preço'
                );
            END IF;
        END LOOP;
    END LOOP;
END $$;