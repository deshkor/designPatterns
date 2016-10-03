public class StrongEnemy implements Enemy {

	StrongEnemy() {
		System.out.println("A strong enemy has been born.");
	}

	private int hp = 50, damage = 10, armor = 10;

	@Override
	public int attack() {
		return this.damage;
	}

	@Override
	public void defend(int damage) {
		if (this.armor >= damage) {
			return;
		}

		this.hp -= (damage - armor);
	}

	@Override
	public int getHp() {
		return this.hp;
	}
}
