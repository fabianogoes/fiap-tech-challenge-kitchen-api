# Localstack

## Running

```shell
docker-compose up -d
```

> Para acessar o dash local, é preciso se local no site no localstack: https://app.localstack.cloud/
> dashboard: https://app.localstack.cloud/dashboard

## AWS CLI configuration

### awslocal instalation ubuntu

```shell
sudo apt update
sudo apt install python3
python3 --version

sudo apt install build-essential zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libreadline-dev libffi-dev wget
sudo apt install -y python3-pip
pip3 --version

pip install awscli-local
awslocal --version
```

```shell
aws configure set aws_access_key_id "dummy" --profile localstack
aws configure set aws_secret_access_key "dummy" --profile localstack
aws configure set region "us-east-1" --profile localstack
aws configure set output "table" --profile localstack
```

## Create SNS

**using awslocal**


```shell
awslocal sns create-topic --name order-payment-events
```

**using aws-cli**

```shell
aws --endpoint-url=http://localhost:4566 \
    sns create-topic --name order-payment-events \
    --region us-east-1 \
    --profile localstack \
    --output table | cat
```

## Create SQS

```shell
awslocal sqs create-queue --queue-name order-payment-queue
```

```shell
aws --endpoint-url=http://localhost:4566 \
  sqs create-queue --queue-name order-payment-queue \
  --region us-east-1 \
  --profile localstack \
  --output table | cat
```

## Subscribe SQS queue to the topic(SNS)

```shell
awslocal --endpoint-url=http://localhost:4566 sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:order-payment-events --protocol sqs --notification-endpoint arn:aws:sqs:us-east-1:000000000000:order-payment-queue
```

```shell
aws --endpoint-url=http://localhost:4566 sns subscribe \
  --topic-arn arn:aws:sns:us-east-1:000000000000:order-payment-events \
  --protocol sqs --notification-endpoint arn:aws:sqs:us-east-1:000000000000:order-payment-queue
```

## References

- [Documentação Oficial](https://docs.localstack.cloud/overview/)
- [Como configurar LocalStack para mapear serviços da AWS](https://medium.com/@valdemarjuniorr/como-configurar-localstack-para-mapear-servicos-da-aws-c8c25e6363b4)
- [LocalStack: simule ambientes AWS localmente](https://www.zup.com.br/blog/localstack-simule-ambientes-aws-localmente)