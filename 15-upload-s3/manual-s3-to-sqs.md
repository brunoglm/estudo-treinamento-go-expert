1 - Crie um bucket S3 e guarde o ARN

2 - Crie uma fila SQS e guarde o ARN

3 - Atualize a politica do SQS para receber eventos do S3
{
  "Version": "2012-10-17",
  "Id": "sqs-policy",
  "Statement": [
    {
      "Sid": "Allow-S3-SendMessage",
      "Effect": "Allow",
      "Principal": {
        "Service": "s3.amazonaws.com"
      },
      "Action": "sqs:SendMessage",
      "Resource": "arn-da-fila",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn-do-s3"
        }
      }
    }
  ]
}

4 - configure o evento no S3:
No console:

S3 → Bucket → Properties → Event Notifications → Create event

Escolha:

Event types:
- PUT
- POST
- Copy
- CompleteMultipartUpload
- Ou só PUT

Prefix: pending/ (opcional — muito útil)
Destination: SQS

Selecione sua fila.
