from django.db import models
from product.models import Product

class BasketItem(models.Model):
    product = models.ForeignKey(Product, on_delete=models.CASCADE)
    basket = models.ForeignKey('Basket', on_delete=models.CASCADE, related_name='items')
    quantity = models.PositiveIntegerField()
    
    @property
    def total_price(self) -> float:
        return float(self.product.price * self.quantity)
    
class Basket(models.Model):
    buyer_id = models.IntegerField()
    order = models.ForeignKey('Order', on_delete=models.CASCADE, null=True, blank=True)
    
    @property
    def total_price(self) -> float:
        return sum([item.total_price for item in self.items.all()])
    
    @property
    def total_items(self) -> int:
        return sum([item.quantity for item in self.items.all()])
    
class Order(models.Model):
    buyer_id = models.IntegerField()
    total_price = models.DecimalField(max_digits=10, decimal_places=2)
    status = models.CharField(max_length=50, choices=[
        ('pending', 'Pending'),
        ('confirmed', 'Confirmed'),
        ('shipped', 'Shipped'),
        ('delivered', 'Delivered'),
        ('cancelled', 'Cancelled'),
    ])
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

