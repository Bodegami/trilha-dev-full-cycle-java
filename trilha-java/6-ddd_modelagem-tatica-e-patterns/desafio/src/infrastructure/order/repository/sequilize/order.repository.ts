import Order from "../../../../domain/checkout/entity/order";
import OrderItem from "../../../../domain/checkout/entity/order_item";
import OrderItemModel from "./order-item.model";
import OrderModel from "./order.model";
import OrderRepositoryInterface from "../../../../domain/checkout/repository/order-repository.interface";

export default class OrderRepository implements OrderRepositoryInterface {
    async create(entity: Order): Promise<void> {
        await OrderModel.create(
            {
                id: entity.id,
                customer_id: entity.customerId,
                total: entity.total(),
                items: entity.items.map((item) => ({
                    id: item.id,
                    name: item.name,
                    price: item.price,
                    product_id: item.productId,
                    quantity: item.quantity,
                })),
            },
            {
                include: [{ model: OrderItemModel }],
            }
        );
    }

    async update(entity: Order): Promise<void> {

        entity.items.map((item) => {
            let orderItem = OrderItemModel.findOne({ where: { id: item.id } });

            if (orderItem) {
                OrderItemModel.update({
                        name: item.name,
                        price: item.price,
                        quantity: item.quantity,
                        productId: item.productId
                    },
                    {
                        where: { id: item.id }
                    });
            }
        });

        await OrderModel.update(
            {
                customer_id: entity.customerId,
                total: entity.total(),
                items: entity.items.map((item) => ({
                    id: item.id,
                    name: item.name,
                    price: item.price,
                    productId: item.productId,
                    quantity: item.quantity
                }))
            },
            {
                where: {
                    id: entity.id,
                },
            },
        )
    }

    async find(id: string): Promise<Order> {
        let orderModel;
        try {
            orderModel = await OrderModel.findOne({
                where: {
                    id,
                },
                include: ["items"],
                rejectOnEmpty: true,
            });
        } catch (error) {
            throw new Error("Order not found");
        }

        const orderItems: OrderItem[] = orderModel.items.map(item => {
            return new OrderItem(item.id, item.name, item.price, item.product_id, item.quantity);
        })

        const order = new Order(id, orderModel.customer_id, orderItems);

        return order;
    }


    async findAll(): Promise<Order[]> {
        const orderModels = await OrderModel.findAll({
            include: ["items"],
        });

        const orders: Order[] = orderModels.map((orderModel) => {
            const orderItems: OrderItem[] = orderModel.items.map(item => {
                return new OrderItem(item.id, item.name, item.price, item.product_id, item.quantity);
            })
            return new Order(orderModel.id, orderModel.customer_id, orderItems);
        });

        return orders;
    }
}
