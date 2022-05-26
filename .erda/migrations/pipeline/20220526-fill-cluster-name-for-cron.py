"""
Generated by Erda Migrator.
Please implement the function entry, and add it to the list entries.
"""


import django.db.models
import json


class PipelineCrons(django.db.models.Model):
    """
    generated by erda-cli
    """

    id = django.db.models.BigIntegerField()
    application_id = django.db.models.BigIntegerField()
    branch = django.db.models.CharField()
    cron_expr = django.db.models.CharField()
    enable = django.db.models.BooleanField()
    pipeline_source = django.db.models.CharField()
    pipeline_yml_name = django.db.models.CharField()
    base_pipeline_id = django.db.models.BigIntegerField()
    extra = django.db.models.TextField()
    time_created = django.db.models.DateTimeField()
    time_updated = django.db.models.DateTimeField()
    pipeline_definition_id = django.db.models.CharField()
    is_edge = django.db.models.BooleanField()
    cluster_name = django.db.models.CharField()
    
    class Meta:
        db_table = "pipeline_crons"


def entry():
    """
    please implement this and add it to the list entries
    """
    crons = PipelineCrons.objects.all()
    for cron in crons:
        if cron.extra is not None:
            clusterName = ""
            try:
                cronExtra = json.loads(cron.extra)
                if "clusterName" in cronExtra:
                    clusterName = cronExtra["clusterName"]
            except Exception as e:
                print("cron extra %s json parse error, exception: %s, skip" % (cluster.name, e))
                continue
            cron.cluster_name = clusterName
            cron.save()


entries: [callable] = [
    entry,
]
