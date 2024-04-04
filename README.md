Brick Backend Test
Simple Transfer API

General flow of the code:
https://mermaid.ink/img/pako:eNrFVU2PmzAQ_SuWz0kKywayqMqp_6C3NqvI4CGxFmzWHyvSKP-9NiYJRJDdHKpyADQez3vzZsY-4lxQwClW8G6A5_CDkZ0k1YbXRGqWs5pwjYwCObToPZMUOcthuJAR_rbhyD5u03y97nmm6LeE99fWB5E8F8Zu-CAlo0QzwdH3TK7Rsf1kwLedx5abKgOZoqZpvIvbv-Wkgqvt5DF7YBba-f17TC40IPEBEnm8FqlDAO_Smm60sLTU66Os7uP3w_f-kdJCgo_Zw7D4ptRjurnKtbqpx3W7z9BHdm-0A40YL4SsegFJJoxGFJQzeXsXebKnfHmpQFoSrgobuUdNCSNzuK3qmd6jBUekcr5frADK95C_DdrhVnZSDjSFpmYSqF-baGcJ59YalOa6x1lvJLotkHcGTj8bmrOk6n9qOjJViHBq1YLcOCeO7NFleqM2cuqoiWRaRlo2W0bTIAh8gppoo9KvFfnSdZYX8ZNGP5upUSYPkuhNkjWyggG9crGS1FLsLJhC31BtK834bsPHlc1JWWbEdqoWfc6e2yWk54NM7fSnk1oP-2Z-jj2p-Tm_Lt175_lFxDbfQ3cpTTTwUOhxHhOYeIYrsKcSo_ZaPDrbBus92LHBqf2lUBA3wHjDT9aVGC1-HniOUy0NzLAXqLtFcVqQUlmrvRpxesQNTl_ixVMSRS9xECTLMAlWM3zA6TyKl4tVEgXJ82qZPMdBfJrhP0LYCOEiWoXLVRKHyygKwyh4asP9ahcd5ukvDvK2EA?type=png

The code using my personal boilerplate that implemented my version of (Rob C Martin) clean architecture which isolating logic layer from any dependencies including data source / repository layer, third party and delivery layer / controller.
Logic located on service folder, data source located on repository folder and represntation layer located on delivery/http folder. Scheme of request and response located on model folder and any call to third party or global utils located on pkg folder.
The application consists 3 endpoints:
- /account-validations
- /transfers
- /transfer-callbacks

Mock of bank api located on:
- https://660f0144356b87a55c50a942.mockapi.io/v1/account-validations
- https://660f0144356b87a55c50a942.mockapi.io/v1/transfers

For quick usage please register this 3 values on your local environment:
- APP_ADDRESS=:8080
- DB_STRING="postgres://avnadmin:AVNS_vyYbIioFIxcp-AVtPhs@pg-3e9f54c9-student-d87c.a.aivencloud.com:23838/defaultdb"
- BANK_URL="https://660f0144356b87a55c50a942.mockapi.io/v1/"
