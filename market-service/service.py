import grpc
import market_pb2
import market_pb2_grpc
from database import SessionLocal
from models import Category, Product, Image


class CategoryService(market_pb2_grpc.CategoryServiceServicer):
    def __init__(self):
        self.db = SessionLocal()

    # Get a Category by ID
    def GetCategory(self, request, context):
        category = self.db.query(Category).filter(Category.id == request.id).first()
        if not category:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("Category not found")
            return market_pb2.CategoryResponse()

        return market_pb2.CategoryResponse(
            id=category.id,
            name=category.name,
            description=category.description,
        )

    # Create a new Category
    def CreateCategory(self, request, context):
        new_category = Category(name=request.name, description=request.description)
        self.db.add(new_category)
        self.db.commit()
        self.db.refresh(new_category)
        return market_pb2.CategoryResponse(
            id=new_category.id,
            name=new_category.name,
            description=new_category.description,
        )

    # List all Categories
    def ListCategories(self, request, context):
        categories = self.db.query(Category).all()
        response = market_pb2.ListCategoriesResponse()
        for category in categories:
            response.categories.add(
                id=category.id,
                name=category.name,
                description=category.description,
            )
        return response

    def UpdateCategory(self, request, context):
        category = self.db.query(Category).filter(Category.id == request.id).first()
        if not category:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("Category not found")
            return market_pb2.CategoryResponse()
        category.name = request.name
        category.description = request.description
        self.db.commit()
        return market_pb2.CategoryResponse(
            id=category.id,
            name=category.name,
            description=category.description,
        )

    def DeleteCategory(self, request, context):
        category = self.db.query(Category).filter(Category.id == request.id).first()
        if not category:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("Category not found")
            return market_pb2.DeleteCategoryResponse()

        self.db.delete(category)
        self.db.commit()
        return market_pb2.DeleteCategoryResponse(id=category.id)


class ProductService(market_pb2_grpc.ProductServiceServicer):
    def __init__(self):
        self.db = SessionLocal()

    # Get a Product by ID
    def GetProduct(self, request, context):
        product = self.db.query(Product).filter(Product.id == request.id).first()
        if not product:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("Product not found")
            return market_pb2.ProductResponse()

        return market_pb2.ProductResponse(
            id=product.id,
            name=product.name,
            description=product.description,
            price=float(product.price),
            quantity=product.quantity,
            rating=product.rating,
            category_id=product.category_id,
        )

    # Create a new Product
    def CreateProduct(self, request, context):
        new_product = Product(
            name=request.name,
            category_id=request.category_id,
            description=request.description,
            price=request.price,
            quantity=request.quantity,
        )
        self.db.add(new_product)
        self.db.commit()
        self.db.refresh(new_product)
        return market_pb2.ProductResponse(
            id=new_product.id,
            name=new_product.name,
            description=new_product.description,
            price=float(new_product.price),
            quantity=new_product.quantity,
            rating=new_product.rating,
            category_id=new_product.category_id,
        )

    # List all Products
    def ListProducts(self, request, context):
        products = self.db.query(Product).all()
        response = market_pb2.ListProductsResponse()
        for product in products:
            response.products.add(
                id=product.id,
                name=product.name,
                description=product.description,
                price=float(product.price),
                quantity=product.quantity,
                rating=product.rating,
                category_id=product.category_id,
            )
        return response

    def UpdateProduct(self, request, context):
        product = self.db.query(Product).filter(Product.id == request.id).first()
        if not product:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("Product not found")
            return market_pb2.ProductResponse()

        product.name = request.name
        product.description = request.description
        product.price = request.price
        product.quantity = request.quantity
        self.db.commit()
        return market_pb2.ProductResponse(
            id=product.id,
            name=product.name,
            description=product.description,
            price=float(product.price),
            quantity=product.quantity,
            category_id=product.category_id,
        )

    def DeleteProduct(self, request, context):
        product = self.db.query(Product).filter(Product.id == request.id).first()
        if not product:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("Product not found")
            return market_pb2.DeleteProductResponse()

        self.db.delete(product)
        self.db.commit()
        return market_pb2.DeleteProductResponse(id=product.id)


class ImageService(market_pb2_grpc.ImageServiceServicer):
    def __init__(self):
        self.db = SessionLocal()

    # List images by product ID
    def ListImagesByProduct(self, request, context):
        images = (
            self.db.query(Image).filter(Image.product_id == request.product_id).all()
        )
        response = market_pb2.ListImagesResponse()
        for image in images:
            response.images.add(
                id=image.id,
                product_id=image.product_id,
                image_url=image.image_url,
                is_primary=image.is_primary,
            )
        return response

    # Add an image to a product
    def AddImage(self, request, context):
        new_image = Image(
            product_id=request.product_id,
            image_url=request.image_url,
            is_primary=request.is_primary,
        )
        self.db.add(new_image)
        self.db.commit()
        self.db.refresh(new_image)
        return market_pb2.ImageResponse(
            id=new_image.id,
            product_id=new_image.product_id,
            image_url=new_image.image_url,
            is_primary=new_image.is_primary,
        )

    # Delete an image
    def DeleteImage(self, request, context):
        image = self.db.query(Image).filter(Image.id == request.id).first()
        if not image:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("Image not found")
            return market_pb2.DeleteImageResponse()

        self.db.delete(image)
        self.db.commit()
        return market_pb2.DeleteImageResponse(id=image.id)
