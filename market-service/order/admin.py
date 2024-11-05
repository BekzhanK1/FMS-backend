from django.contrib import admin
from .models import Order, Basket, BasketItem

admin.site.register(Order)
admin.site.register(Basket)
admin.site.register(BasketItem)

# Register your models here.
