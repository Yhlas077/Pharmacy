// Declarative config for each CRUD resource. Drives the generic CrudView:
// columns (table), fields (form), filters (toolbar), and which api module to call.

import { users, categories, pharmacies, medicines } from './api/index.js'
import { ROLES } from './api/constants.js'

const money = (v) => (v ? '₼' + Number(v).toFixed(2) : '—')

export const resources = {
  users: {
    title: 'Users',
    singular: 'User',
    api: users,
    columns: [
      { key: 'id', label: 'ID', width: '60px' },
      { key: 'name', label: 'Name' },
      { key: 'phone', label: 'Phone' },
      { key: 'region', label: 'Region' },
      { key: 'email', label: 'Email' },
      { key: 'role', label: 'Role', badge: true },
    ],
    filters: [
      { key: 'search', type: 'search', placeholder: 'Search name, email, phone…' },
      { key: 'role', type: 'select', label: 'Role', options: ROLES, allLabel: 'All roles' },
    ],
    fields: [
      { key: 'name', label: 'Name', required: true },
      { key: 'phone', label: 'Phone', required: true },
      { key: 'region', label: 'Region', required: true },
      { key: 'email', label: 'Email', type: 'email', required: true },
      { key: 'password', label: 'Password', type: 'password', required: true, createOnly: true },
      { key: 'role', label: 'Role', type: 'select', options: ROLES, default: 'user', required: true },
    ],
  },

  categories: {
    title: 'Categories',
    singular: 'Category',
    api: categories,
    columns: [
      { key: 'id', label: 'ID', width: '60px' },
      { key: 'name', label: 'Name' },
    ],
    filters: [{ key: 'search', type: 'search', placeholder: 'Search categories…' }],
    fields: [{ key: 'name', label: 'Name', required: true }],
  },

  medicines: {
    title: 'Medicines',
    singular: 'Medicine',
    api: medicines,
    columns: [
      { key: 'id', label: 'ID', width: '60px' },
      { key: 'name', label: 'Name' },
      { key: 'price', label: 'Price', format: money },
      { key: 'new_price', label: 'New price', format: money },
      { key: 'stock', label: 'Stock', stock: true },
      { key: 'category_id', label: 'Category', lookup: 'categories' },
      { key: 'pharmacy_id', label: 'Pharmacy', lookup: 'pharmacies' },
    ],
    filters: [
      { key: 'search', type: 'search', placeholder: 'Search medicines…' },
      { key: 'category_id', type: 'select', label: 'Category', lookup: 'categories', allLabel: 'All categories' },
    ],
    fields: [
      { key: 'name', label: 'Name', required: true },
      { key: 'description', label: 'Description', type: 'textarea' },
      { key: 'price', label: 'Price', type: 'number', required: true },
      { key: 'new_price', label: 'New price (discount)', type: 'number' },
      { key: 'stock', label: 'Stock', type: 'number', required: true, default: 0 },
      { key: 'category_id', label: 'Category', type: 'select', lookup: 'categories', required: true },
      { key: 'pharmacy_id', label: 'Pharmacy', type: 'select', lookup: 'pharmacies', required: true },
    ],
  },

  pharmacies: {
    title: 'Pharmacies',
    singular: 'Pharmacy',
    api: pharmacies,
    columns: [
      { key: 'id', label: 'ID', width: '60px' },
      { key: 'name', label: 'Name' },
      { key: 'address', label: 'Address' },
      { key: 'pharmacy_hours', label: 'Hours' },
    ],
    filters: [{ key: 'search', type: 'search', placeholder: 'Search pharmacies…' }],
    fields: [
      { key: 'name', label: 'Name', required: true },
      { key: 'address', label: 'Address', required: true },
      { key: 'pharmacy_hours', label: 'Hours', placeholder: 'e.g. 08:00–22:00 or 24/7', required: true },
    ],
  },
}
