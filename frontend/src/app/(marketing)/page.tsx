import { Hero } from "@/features/marketing/components/Hero";

export default function MarketingPage() {
  return (
    <main className="min-h-screen bg-background">
      <header className="container mx-auto py-6 px-4 flex justify-between items-center">
        <div className="font-bold text-2xl tracking-tighter">SubWatch</div>
        <nav className="text-sm font-medium text-muted-foreground hidden md:flex gap-6">
            <a href="#" className="hover:text-foreground transition-colors">Funcionalidades</a>
            <a href="#" className="hover:text-foreground transition-colors">Preço</a>
            <a href="#" className="hover:text-foreground transition-colors">Sobre</a>
        </nav>
      </header>
      <Hero />
    </main>
  );
}
