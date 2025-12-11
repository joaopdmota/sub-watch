import App from './app';

try {
    const app = new App();
    app.start();
} catch (error) {
    console.error('Error on app start:', error);
    process.exit(1);
}