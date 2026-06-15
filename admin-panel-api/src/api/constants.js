// Shared enums (previously in mockData.js).
export const ORDER_STATUSES = ['pending', 'paid', 'shipped', 'delivered', 'cancelled']
export const ROLES = ['admin', 'pharmacy', 'user']

// Allowed order status transitions (mirrors the backend rules in API.md).
export const ORDER_TRANSITIONS = {
  pending: ['paid', 'cancelled'],
  paid: ['shipped', 'cancelled'],
  shipped: ['delivered'],
  delivered: [],
  cancelled: [],
}
