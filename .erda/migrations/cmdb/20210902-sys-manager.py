"""
Generated by Erda Migrator.
Please implement the function entry, and add it to the list entries.
"""


import django.db.models
from datetime import datetime


class DiceMemberExtra(django.db.models.Model):
    """
    generated by erda-cli
    """

    id = django.db.models.BigIntegerField()
    created_at = django.db.models.DateTimeField()
    updated_at = django.db.models.DateTimeField()
    user_id = django.db.models.CharField()
    parent_id = django.db.models.CharField()
    scope_id = django.db.models.CharField()
    scope_type = django.db.models.CharField()
    resource_key = django.db.models.CharField()
    resource_value = django.db.models.CharField()

    class Meta:
        db_table = "dice_member_extra"


def entry():
    """
    please implement this and add it to the list entries
    """
    DiceMemberExtra.objects.filter(
        user_id="1", scope_type="sys", scope_id="0").delete()
    m = DiceMemberExtra()
    m.created_at = datetime.now()
    m.updated_at = datetime.now()
    m.user_id = '1'
    m.parent_id = '0'
    m.scope_id = '0'
    m.scope_type = 'sys'
    m.resource_key = 'role'
    m.resource_value = 'Manager'
    m.save()
    pass


entries: [callable] = [
    entry,
]
