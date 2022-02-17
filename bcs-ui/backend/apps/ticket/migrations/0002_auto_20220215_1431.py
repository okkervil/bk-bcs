# Generated by Django 3.2.2 on 2022-02-15 06:31

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('ticket', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='tlscert',
            name='cert',
            field=models.TextField(blank=True, null=True, verbose_name='证书内容'),
        ),
        migrations.AlterField(
            model_name='tlscert',
            name='creator',
            field=models.CharField(max_length=64, verbose_name='创建者'),
        ),
        migrations.AlterField(
            model_name='tlscert',
            name='key',
            field=models.TextField(blank=True, null=True, verbose_name='证书内容'),
        ),
        migrations.AlterField(
            model_name='tlscert',
            name='updator',
            field=models.CharField(max_length=64, verbose_name='修改者'),
        ),
    ]
