import Link from "next/link";
import { Button } from "@/components/ui";

export function Hero() {
  return (
    <section className="text-center py-20 px-4">
      <h1 className="text-4xl md:text-6xl font-extrabold mb-6 bg-gradient-to-r from-white to-gray-400 bg-clip-text text-transparent">
        Domine suas assinaturas.
      </h1>
      <p className="text-xl text-muted-foreground max-w-2xl mx-auto mb-10">
        Gerencie, economize e visualize todos os seus gastos recorrentes em um único lugar.
        Simples, bonito e eficiente.
      </p>
      <div className="flex gap-4 justify-center">
        <Link href="/dashboard">
            <Button className="h-12 px-8 text-base">Começar Agora</Button>
        </Link>
        <Button className="h-12 px-8 text-base bg-secondary text-secondary-foreground hover:bg-secondary/80">
          Saber Mais
        </Button>
      </div>
    </section>
  );
}
