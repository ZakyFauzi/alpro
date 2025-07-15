// Data produk jam tangan
const products = [
    {
        id: 1,
        name: "Rolex Submariner",
        category: "luxury",
        price: 85000000,
        image: "https://images.unsplash.com/photo-1524592094714-0f0654e20314?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80",
        badge: "Best Seller",
        description: "Jam tangan mewah dengan desain klasik dan tahan air hingga 300m"
    },
    {
        id: 2,
        name: "Omega Speedmaster",
        category: "sport",
        price: 45000000,
        image: "https://images.unsplash.com/photo-1558618666-fcd25c85cd64?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80",
        badge: "Limited",
        description: "Jam tangan sport dengan chronograph dan tahan goncangan"
    },
    {
        id: 3,
        name: "Cartier Tank",
        category: "business",
        price: 35000000,
        image: "https://images.unsplash.com/photo-1547996160-81dfa63595aa?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80",
        badge: "Elegant",
        description: "Jam tangan formal dengan desain persegi yang elegan"
    },
    {
        id: 4,
        name: "Seiko 5 Sports",
        category: "casual",
        price: 2500000,
        image: "https://images.unsplash.com/photo-1523275335684-37898b6baf30?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80",
        badge: "Popular",
        description: "Jam tangan casual yang nyaman untuk sehari-hari"
    },
    {
        id: 5,
        name: "Patek Philippe Nautilus",
        category: "luxury",
        price: 120000000,
        image: "https://images.unsplash.com/photo-1558618666-fcd25c85cd64?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80",
        badge: "Premium",
        description: "Jam tangan premium dengan mekanisme otomatis"
    },
    {
        id: 6,
        name: "Garmin Fenix",
        category: "sport",
        price: 8500000,
        image: "https://images.unsplash.com/photo-1523275335684-37898b6baf30?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80",
        badge: "Smart",
        description: "Smartwatch dengan GPS dan fitur olahraga lengkap"
    },
    {
        id: 7,
        name: "Tissot Le Locle",
        category: "business",
        price: 8500000,
        image: "https://images.unsplash.com/photo-1547996160-81dfa63595aa?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80",
        badge: "Classic",
        description: "Jam tangan bisnis dengan desain klasik Swiss"
    },
    {
        id: 8,
        name: "Casio G-Shock",
        category: "casual",
        price: 1500000,
        image: "https://images.unsplash.com/photo-1523275335684-37898b6baf30?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80",
        badge: "Durable",
        description: "Jam tangan tahan banting untuk aktivitas outdoor"
    }
];

// State aplikasi
let cart = [];
let currentFilter = 'all';

// DOM Elements
const productsGrid = document.getElementById('productsGrid');
const filterButtons = document.querySelectorAll('.filter-btn');
const cartIcon = document.querySelector('.cart-icon');
const cartModal = document.getElementById('cartModal');
const closeCart = document.getElementById('closeCart');
const cartItems = document.getElementById('cartItems');
const cartTotal = document.getElementById('cartTotal');
const cartCount = document.querySelector('.cart-count');
const checkoutBtn = document.getElementById('checkoutBtn');
const hamburger = document.querySelector('.hamburger');
const navMenu = document.querySelector('.nav-menu');
const searchInput = document.querySelector('.search-box input');

// Format currency
function formatCurrency(amount) {
    return new Intl.NumberFormat('id-ID', {
        style: 'currency',
        currency: 'IDR',
        minimumFractionDigits: 0
    }).format(amount);
}

// Render produk
function renderProducts(filteredProducts = products) {
    productsGrid.innerHTML = filteredProducts.map(product => `
        <div class="product-card" data-category="${product.category}">
            <div class="product-image">
                <img src="${product.image}" alt="${product.name}">
                ${product.badge ? `<div class="product-badge">${product.badge}</div>` : ''}
            </div>
            <div class="product-info">
                <div class="product-category">${product.category.toUpperCase()}</div>
                <h3 class="product-title">${product.name}</h3>
                <div class="product-price">${formatCurrency(product.price)}</div>
                <p>${product.description}</p>
                <div class="product-actions">
                    <button class="add-to-cart" onclick="addToCart(${product.id})">
                        <i class="fas fa-shopping-cart"></i> Tambah ke Keranjang
                    </button>
                    <button class="wishlist" onclick="addToWishlist(${product.id})">
                        <i class="fas fa-heart"></i>
                    </button>
                </div>
            </div>
        </div>
    `).join('');
}

// Filter produk
function filterProducts(category) {
    currentFilter = category;
    const filteredProducts = category === 'all' 
        ? products 
        : products.filter(product => product.category === category);
    renderProducts(filteredProducts);
}

// Tambah ke keranjang
function addToCart(productId) {
    const product = products.find(p => p.id === productId);
    const existingItem = cart.find(item => item.id === productId);
    
    if (existingItem) {
        existingItem.quantity += 1;
    } else {
        cart.push({
            ...product,
            quantity: 1
        });
    }
    
    updateCart();
    showNotification('Produk berhasil ditambahkan ke keranjang!');
}

// Update keranjang
function updateCart() {
    // Update cart count
    const totalItems = cart.reduce((sum, item) => sum + item.quantity, 0);
    cartCount.textContent = totalItems;
    
    // Update cart items
    cartItems.innerHTML = cart.map(item => `
        <div class="cart-item">
            <div class="cart-item-image">
                <img src="${item.image}" alt="${item.name}">
            </div>
            <div class="cart-item-details">
                <div class="cart-item-title">${item.name}</div>
                <div class="cart-item-price">${formatCurrency(item.price)}</div>
                <div class="cart-item-quantity">
                    <button class="quantity-btn" onclick="updateQuantity(${item.id}, -1)">-</button>
                    <span>${item.quantity}</span>
                    <button class="quantity-btn" onclick="updateQuantity(${item.id}, 1)">+</button>
                </div>
                <button class="remove-item" onclick="removeFromCart(${item.id})">Hapus</button>
            </div>
        </div>
    `).join('');
    
    // Update total
    const total = cart.reduce((sum, item) => sum + (item.price * item.quantity), 0);
    cartTotal.textContent = formatCurrency(total);
}

// Update quantity
function updateQuantity(productId, change) {
    const item = cart.find(item => item.id === productId);
    if (item) {
        item.quantity += change;
        if (item.quantity <= 0) {
            removeFromCart(productId);
        } else {
            updateCart();
        }
    }
}

// Hapus dari keranjang
function removeFromCart(productId) {
    cart = cart.filter(item => item.id !== productId);
    updateCart();
}

// Toggle cart modal
function toggleCart() {
    cartModal.classList.toggle('active');
}

// Checkout
function checkout() {
    if (cart.length === 0) {
        showNotification('Keranjang belanja kosong!', 'error');
        return;
    }
    
    const total = cart.reduce((sum, item) => sum + (item.price * item.quantity), 0);
    alert(`Total pembayaran: ${formatCurrency(total)}\n\nTerima kasih telah berbelanja di Luxury Watches!`);
    
    // Reset cart
    cart = [];
    updateCart();
    toggleCart();
}

// Add to wishlist (placeholder)
function addToWishlist(productId) {
    showNotification('Produk ditambahkan ke wishlist!');
}

// Show notification
function showNotification(message, type = 'success') {
    const notification = document.createElement('div');
    notification.className = `notification ${type}`;
    notification.textContent = message;
    notification.style.cssText = `
        position: fixed;
        top: 100px;
        right: 20px;
        background: ${type === 'success' ? '#27ae60' : '#e74c3c'};
        color: white;
        padding: 1rem 2rem;
        border-radius: 8px;
        z-index: 3000;
        animation: slideIn 0.3s ease;
    `;
    
    document.body.appendChild(notification);
    
    setTimeout(() => {
        notification.style.animation = 'slideOut 0.3s ease';
        setTimeout(() => {
            document.body.removeChild(notification);
        }, 300);
    }, 3000);
}

// Search functionality
function searchProducts(query) {
    const filteredProducts = products.filter(product => 
        product.name.toLowerCase().includes(query.toLowerCase()) ||
        product.category.toLowerCase().includes(query.toLowerCase()) ||
        product.description.toLowerCase().includes(query.toLowerCase())
    );
    renderProducts(filteredProducts);
}

// Smooth scrolling untuk navigasi
function smoothScroll(target) {
    const element = document.querySelector(target);
    if (element) {
        element.scrollIntoView({
            behavior: 'smooth',
            block: 'start'
        });
    }
}

// Event Listeners
document.addEventListener('DOMContentLoaded', function() {
    // Render produk awal
    renderProducts();
    
    // Filter buttons
    filterButtons.forEach(button => {
        button.addEventListener('click', () => {
            // Remove active class from all buttons
            filterButtons.forEach(btn => btn.classList.remove('active'));
            // Add active class to clicked button
            button.classList.add('active');
            // Filter products
            filterProducts(button.dataset.filter);
        });
    });
    
    // Cart functionality
    cartIcon.addEventListener('click', toggleCart);
    closeCart.addEventListener('click', toggleCart);
    checkoutBtn.addEventListener('click', checkout);
    
    // Close cart when clicking outside
    cartModal.addEventListener('click', (e) => {
        if (e.target === cartModal) {
            toggleCart();
        }
    });
    
    // Mobile menu
    hamburger.addEventListener('click', () => {
        navMenu.classList.toggle('active');
        hamburger.classList.toggle('active');
    });
    
    // Close mobile menu when clicking on a link
    document.querySelectorAll('.nav-link').forEach(link => {
        link.addEventListener('click', () => {
            navMenu.classList.remove('active');
            hamburger.classList.remove('active');
        });
    });
    
    // Search functionality
    searchInput.addEventListener('input', (e) => {
        const query = e.target.value.trim();
        if (query === '') {
            renderProducts();
        } else {
            searchProducts(query);
        }
    });
    
    // Smooth scrolling for navigation links
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function (e) {
            e.preventDefault();
            const target = this.getAttribute('href');
            smoothScroll(target);
        });
    });
    
    // Contact form
    const contactForm = document.querySelector('.contact-form');
    contactForm.addEventListener('submit', (e) => {
        e.preventDefault();
        showNotification('Pesan berhasil dikirim! Kami akan menghubungi Anda segera.');
        contactForm.reset();
    });
    
    // Add CSS animations
    const style = document.createElement('style');
    style.textContent = `
        @keyframes slideIn {
            from {
                transform: translateX(100%);
                opacity: 0;
            }
            to {
                transform: translateX(0);
                opacity: 1;
            }
        }
        
        @keyframes slideOut {
            from {
                transform: translateX(0);
                opacity: 1;
            }
            to {
                transform: translateX(100%);
                opacity: 0;
            }
        }
        
        .hamburger.active span:nth-child(1) {
            transform: rotate(-45deg) translate(-5px, 6px);
        }
        
        .hamburger.active span:nth-child(2) {
            opacity: 0;
        }
        
        .hamburger.active span:nth-child(3) {
            transform: rotate(45deg) translate(-5px, -6px);
        }
    `;
    document.head.appendChild(style);
    
    // Intersection Observer untuk animasi scroll
    const observerOptions = {
        threshold: 0.1,
        rootMargin: '0px 0px -50px 0px'
    };
    
    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.style.opacity = '1';
                entry.target.style.transform = 'translateY(0)';
            }
        });
    }, observerOptions);
    
    // Observe elements for animation
    document.querySelectorAll('.product-card, .category-card, .feature').forEach(el => {
        el.style.opacity = '0';
        el.style.transform = 'translateY(30px)';
        el.style.transition = 'opacity 0.6s ease, transform 0.6s ease';
        observer.observe(el);
    });
});

// Keyboard shortcuts
document.addEventListener('keydown', (e) => {
    // ESC to close cart
    if (e.key === 'Escape' && cartModal.classList.contains('active')) {
        toggleCart();
    }
    
    // Ctrl/Cmd + K to focus search
    if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
        e.preventDefault();
        searchInput.focus();
    }
});

// Lazy loading untuk gambar
function lazyLoadImages() {
    const images = document.querySelectorAll('img[data-src]');
    const imageObserver = new IntersectionObserver((entries, observer) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                const img = entry.target;
                img.src = img.dataset.src;
                img.removeAttribute('data-src');
                imageObserver.unobserve(img);
            }
        });
    });
    
    images.forEach(img => imageObserver.observe(img));
}

// Initialize lazy loading
lazyLoadImages();