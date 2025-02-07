CREATE TABLE IF NOT EXISTS constantin_kanivec
(
    id                   INT AUTO_INCREMENT     PRIMARY KEY,
    order_id             INT                    null COMMENT 'Уникальный идентификатор заказа',
    customer_id          INT                    null COMMENT 'Уникальный идентификатор клиента',
    order_number         VARCHAR(20)            null COMMENT 'Произвольный номер заказа',
    fakt_data            DATE                   null COMMENT 'Зyачение произвольного поля в CRM-системе',
    podtverzden_otpravka DATE                   null COMMENT 'Значение произвольного поля в CRM-системе',
    total_summ           DOUBLE                 null COMMENT 'Сумма заказа',
    prepay_sum           DOUBLE                 null COMMENT 'Оплаченная сумма'

)
COMMENT 'Не забыть заменить на ссылку на репозиторий';
